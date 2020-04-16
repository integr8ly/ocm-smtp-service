package logging

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/logger"
	"net/http"
)

func NewLoggingWriter(w http.ResponseWriter, r *http.Request, f LogFormatter) *loggingWriter {
	return &loggingWriter{ResponseWriter: w, request: r, formatter: f}
}

type loggingWriter struct {
	http.ResponseWriter
	request        *http.Request
	formatter      LogFormatter
	responseStatus int
	responseBody   []byte
}

func (writer *loggingWriter) Write(body []byte) (int, error) {
	writer.responseBody = body
	return writer.ResponseWriter.Write(body)
}

func (writer *loggingWriter) WriteHeader(status int) {
	writer.responseStatus = status
	writer.ResponseWriter.WriteHeader(status)
}

func (writer *loggingWriter) log(log string, err error) {
	olog := logger.NewOCMLogger(writer.request.Context())
	switch err {
	case nil:
		olog.V(LoggingThreshold).Infof(log)
	default:
		olog.Errorf("Unable to format request/response for log. Error: %s", err.Error())
	}
}

func (writer *loggingWriter) prepareRequestLog() (string, error) {
	return writer.formatter.FormatRequestLog(writer.request)
}

func (writer *loggingWriter) prepareResponseLog(elapsed string) (string, error) {
	info := &ResponseInfo{
		Header:  writer.ResponseWriter.Header(),
		Body:    writer.responseBody,
		Status:  writer.responseStatus,
		Elapsed: elapsed,
	}

	return writer.formatter.FormatResponseLog(info)
}

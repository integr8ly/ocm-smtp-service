FROM centos:7

RUN \
    yum install -y \
    openssl \
    postgresql \
    && \
    yum clean all

COPY \
    ocm-smtp-service \
    /usr/local/bin/

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/ocm-smtp-service", "serve"]

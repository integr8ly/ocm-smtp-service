package db

import (
	"context"
	"database/sql"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/logger"
)

type contextKey int

const (
	transactionKey contextKey = iota
)

// NewContext returns a new context with transaction stored in it.
// Upon error, the original context is still returned along with an error
func NewContext(ctx context.Context) (context.Context, error) {
	tx, err := newTransaction()
	if err != nil {
		return ctx, err
	}

	// adding txid explicitly to context with a simple string key and int value
	// due to a cyclical import cycle between pkg/db and pkg/logging
	ctx = context.WithValue(ctx, "txid", tx.txid)
	ctx = context.WithValue(ctx, transactionKey, tx)

	return ctx, nil
}

// TxContext creates a new transaction context from context.Background()
func TxContext() (ctx context.Context, err error) {
	return NewContext(context.Background())
}

// Resolve resolves the current transaction according to the rollback flag.
func Resolve(ctx context.Context) {
	olog := logger.NewOCMLogger(ctx)
	tx, ok := ctx.Value(transactionKey).(*txFactory)
	if !ok {
		olog.Errorf("Could not retrieve transaction from context")
		return
	}

	if tx.markedForRollback() {
		if err := rollback(ctx); err != nil {
			olog.Errorf("Could not rollback transaction: %v", err)
		}
		olog.Infof("Rolled back transaction")
	} else {
		if err := commit(ctx); err != nil {
			// TODO:  what does the user see when this occurs? seems like they will get a false positive
			olog.Errorf("Could not commit transaction: %v", err)
			return
		}
	}
}

// FromContext Retrieves the transaction from the context.
func FromContext(ctx context.Context) (*sql.Tx, error) {
	transaction, ok := ctx.Value(transactionKey).(*txFactory)
	if !ok {
		return nil, errors.GeneralError("Could not retrieve transaction from context")
	}
	return transaction.tx, nil
}

// MarkForRollback flags the transaction stored in the context for rollback and logs whatever error caused the rollback
func MarkForRollback(ctx context.Context, err error) {
	olog := logger.NewOCMLogger(ctx)
	transaction, ok := ctx.Value(transactionKey).(*txFactory)
	if !ok {
		olog.Errorf("failed to mark transaction for rollback: could not retrieve transaction from context")
		return
	}
	olog.Infof("Marked transaction for rollback, err: %v", err)
	transaction.rollbackFlag.val = true
}

// commit commits the transaction stored in context or returns an err if one occurred.
func commit(ctx context.Context) error {
	tx, err := FromContext(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// rollback rollbacks the transaction stored in context or returns an err if one occurred..
func rollback(ctx context.Context) error {
	tx, err := FromContext(ctx)
	if err != nil {
		return err
	}
	return tx.Rollback()
}

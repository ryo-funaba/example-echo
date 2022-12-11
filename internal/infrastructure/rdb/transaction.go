package rdb

import (
	"context"

	"github.com/ryo-funaba/example-serverless-go/pkg/errorutil"
)

const transactionStructName = "Transaction"

type transaction struct {
	rdb Cluster
}

type Transaction interface {
	ExecTx(
		ctx context.Context,
		f func(ctx context.Context) error,
		database string,
	) error
}

func NewTransaction(isDebugMode bool) Transaction {
	c := NewCluster(isDebugMode)

	return &transaction{c}
}

func (r *transaction) ExecTx(
		ctx context.Context,
		f func(ctx context.Context) error,
		database string,
	) error {
	con, err := r.rdb.GetConnection(Primary, database)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	tx, err := con.BeginTx(ctx, nil)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	c := r.rdb.SetTxInCtx(ctx, tx)

	err = f(c)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errorutil.Errorf(errorutil.Unknown, err.Error())
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errorutil.Errorf(errorutil.Unknown, err.Error())
		}

		return err
	}

	return nil
}

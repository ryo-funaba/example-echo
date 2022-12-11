package rdb

import (
	"context"
	"database/sql"

	"github.com/ryo-funaba/example-serverless-go/pkg/errorutil"
	"github.com/volatiletech/sqlboiler/boil"
)

const txCtxKey = "transaction_context_key"

type cluster struct {
	Pools       map[string]*sql.DB
	isDebugMode bool
}

type Cluster interface {
	GetConnection(conn ConnectionType, database string) (*sql.DB, error)
	AddConnection(conn ConnectionType, database, driver, dns string) error
	AddConnectionForTest(conn ConnectionType, database, driver, dns string) error
	GetBoilCtxExecutor(ctx context.Context, conn ConnectionType, database string) (boil.ContextExecutor, error)
	SetTxInCtx(ctx context.Context, tx *sql.Tx) context.Context
}

func NewCluster(isDebugMode bool) Cluster {
	pools := map[string]*sql.DB{}
	c := &cluster{pools, isDebugMode}

	return c
}

func (c *cluster) GetConnection(conn ConnectionType, database string) (*sql.DB, error) {
	con, ok := c.Pools[database]
	if !ok {
		return nil, errorutil.Errorf(errorutil.Unknown, "get a connection from pools failed")
	}

	return con, nil
}

func (c *cluster) AddConnection(conn ConnectionType, database, driver, dns string) error {
	con, err := sql.Open(driver, dns)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	err = con.Ping()
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	boil.SetDB(con)
	boil.DebugMode = c.isDebugMode

	c.Pools[database] = con

	return nil
}

func (c *cluster) AddConnectionForTest(conn ConnectionType, database, driver, dns string) error {
	con, err := sql.Open(driver, dns)
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	err = con.Ping()
	if err != nil {
		return errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	boil.SetDB(con)
	boil.DebugMode = c.isDebugMode

	c.Pools[database] = con

	return nil
}

func (c *cluster) GetBoilCtxExecutor(ctx context.Context, conn ConnectionType, database string) (boil.ContextExecutor, error) {
	tx, ok := ctx.Value(txCtxKey).(*sql.Tx)
	if !ok {
		con, err := c.GetConnection(conn, database)
		if err != nil {
			return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
		}

		return con, nil
	}

	return tx, nil
}

func (c *cluster) SetTxInCtx(ctx context.Context, tx *sql.Tx) context.Context {
	v := context.WithValue(ctx, txCtxKey, tx)

	return v
}

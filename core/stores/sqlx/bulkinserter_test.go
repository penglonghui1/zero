package sqlx

import (
	"context"
	"database/sql"

	"github.com/pengcainiao/sqlx"
)

type mockedConn struct {
	query   string
	args    []interface{}
	execErr error
}

func (c *mockedConn) ExecCtx(_ context.Context, query string, args ...interface{}) (sql.Result, error) {
	c.query = query
	c.args = args
	return nil, c.execErr
}

func (c *mockedConn) PrepareCtx(ctx context.Context, query string) (StmtSession, error) {
	panic("implement me")
}

func (c *mockedConn) QueryRowCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (c *mockedConn) QueryRowPartialCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (c *mockedConn) QueryRowsCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (c *mockedConn) QueryRowsPartialCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (c *mockedConn) TransactCtx(ctx context.Context, fn func(context.Context, Session) error) error {
	panic("should not called")
}

func (c *mockedConn) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.ExecCtx(context.Background(), query, args...)
}

func (c *mockedConn) Prepare(query string) (StmtSession, error) {
	panic("should not called")
}

func (c *mockedConn) QueryRow(v interface{}, query string, args ...interface{}) error {
	panic("should not called")
}

func (c *mockedConn) QueryRowPartial(v interface{}, query string, args ...interface{}) error {
	panic("should not called")
}

func (c *mockedConn) QueryRows(v interface{}, query string, args ...interface{}) error {
	panic("should not called")
}

func (c *mockedConn) QueryRowsPartial(v interface{}, query string, args ...interface{}) error {
	panic("should not called")
}

func (c *mockedConn) RawDB() (*sqlx.DB, error) {
	panic("should not called")
}

func (c *mockedConn) Transact(func(session Session) error) error {
	panic("should not called")
}

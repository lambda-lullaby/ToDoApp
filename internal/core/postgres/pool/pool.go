package pool

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNoRows               = errors.New("no rows in result set")
	ErrViolatesForeignKey   = errors.New("violates foreign key constraint")
	ErrViolatesUniqueConstr = errors.New("violates unique constraint")
)

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
	Close()
}

type CommandTag interface {
	RowsAffected() int64
}

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Exec(ctx context.Context, sql string, args ...any) (CommandTag, error)
	Ping(ctx context.Context) error
	Close()

	OpTimeout() time.Duration
}

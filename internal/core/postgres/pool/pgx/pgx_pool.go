package pgx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	core_postgres_pool "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool"
)

const (
	pgErrCodeForeignKeyViolation = "23503"
	pgErrCodeUniqueViolation     = "23505"
)

type Pool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func New(ctx context.Context, dsn string, opTimeout time.Duration) (*Pool, error) {
	pgxPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("create pgx pool: %w", err)
	}

	p := &Pool{Pool: pgxPool, opTimeout: opTimeout}
	if err := p.Ping(ctx); err != nil {
		pgxPool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}
	return p, nil
}

func (p *Pool) OpTimeout() time.Duration {
	return p.opTimeout
}

func (p *Pool) Query(ctx context.Context, sql string, args ...any) (core_postgres_pool.Rows, error) {
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, mapErr(err)
	}
	return &pgxRows{rows: rows}, nil
}

func (p *Pool) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	return &pgxRow{row: p.Pool.QueryRow(ctx, sql, args...)}
}

func (p *Pool) Exec(ctx context.Context, sql string, args ...any) (core_postgres_pool.CommandTag, error) {
	tag, err := p.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return nil, mapErr(err)
	}
	return pgxCommandTag{tag: tag}, nil
}

type pgxRow struct {
	row pgx.Row
}

func (r *pgxRow) Scan(dest ...any) error {
	if err := r.row.Scan(dest...); err != nil {
		return mapErr(err)
	}
	return nil
}

type pgxRows struct {
	rows pgx.Rows
}

func (r *pgxRows) Next() bool {
	return r.rows.Next()
}

func (r *pgxRows) Scan(dest ...any) error {
	if err := r.rows.Scan(dest...); err != nil {
		return mapErr(err)
	}
	return nil
}

func (r *pgxRows) Err() error {
	if err := r.rows.Err(); err != nil {
		return mapErr(err)
	}
	return nil
}

func (r *pgxRows) Close() {
	r.rows.Close()
}

type pgxCommandTag struct {
	tag pgconn.CommandTag
}

func (t pgxCommandTag) RowsAffected() int64 {
	return t.tag.RowsAffected()
}

func mapErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%w", core_postgres_pool.ErrNoRows)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgErrCodeForeignKeyViolation:
			return fmt.Errorf("%s: %w", pgErr.Message, core_postgres_pool.ErrViolatesForeignKey)
		case pgErrCodeUniqueViolation:
			return fmt.Errorf("%s: %w", pgErr.Message, core_postgres_pool.ErrViolatesUniqueConstr)
		}
	}

	return err
}

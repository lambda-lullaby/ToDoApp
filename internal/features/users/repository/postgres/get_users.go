package users_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var query strings.Builder
	var args []any

	query.WriteString(`SELECT id, version, full_name, phone_number FROM todoapp.users ORDER BY id`)
	if limit != nil {
		args = append(args, *limit)
		fmt.Fprintf(&query, " LIMIT $%d", len(args))
	}
	if offset != nil {
		args = append(args, *offset)
		fmt.Fprintf(&query, " OFFSET $%d", len(args))
	}
	query.WriteString(";")

	rows, err := r.pool.Query(ctx, query.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var m UserModel
		if err := m.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		users = append(users, modelToDomain(m))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}

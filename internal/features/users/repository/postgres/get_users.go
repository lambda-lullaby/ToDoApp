package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, version, full_name, phone_number FROM todoapp.users ORDER BY id`

	var args []any
	if limit != nil {
		args = append(args, *limit)
		query += fmt.Sprintf(" LIMIT $%d", len(args))
	}
	if offset != nil {
		args = append(args, *offset)
		query += fmt.Sprintf(" OFFSET $%d", len(args))
	}
	query += ";"

	rows, err := r.pool.Query(ctx, query, args...)
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

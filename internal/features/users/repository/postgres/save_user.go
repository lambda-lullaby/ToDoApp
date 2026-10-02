package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

func (r *UsersRepository) SaveUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.users (id, version, full_name, phone_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, full_name, phone_number;`

	row := r.pool.QueryRow(ctx, query, user.ID, user.Version, user.FullName, user.PhoneNumber)

	var m UserModel
	if err := m.Scan(row); err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}
	return modelToDomain(m), nil
}

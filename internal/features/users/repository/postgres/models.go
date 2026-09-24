package users_postgres_repository

import (
	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type UserModel struct {
	ID          uuid.UUID
	Version     int64
	FullName    string
	PhoneNumber *string
}

type scanner interface {
	Scan(dest ...any) error
}

func (m *UserModel) Scan(row scanner) error {
	return row.Scan(&m.ID, &m.Version, &m.FullName, &m.PhoneNumber)
}

func modelToDomain(m UserModel) domain.User {
	return domain.User{
		ID:          m.ID,
		Version:     m.Version,
		FullName:    m.FullName,
		PhoneNumber: m.PhoneNumber,
	}
}

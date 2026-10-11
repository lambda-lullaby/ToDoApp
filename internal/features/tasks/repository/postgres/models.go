package tasks_postgres_repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type TaskModel struct {
	ID           uuid.UUID
	Version      int64
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID uuid.UUID
}

type scanner interface {
	Scan(dest ...any) error
}

func (m *TaskModel) Scan(row scanner) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.Title,
		&m.Description,
		&m.Completed,
		&m.CreatedAt,
		&m.CompletedAt,
		&m.AuthorUserID,
	)
}

func modelToDomain(m TaskModel) domain.Task {
	return domain.Task{
		ID:           m.ID,
		Version:      m.Version,
		Title:        m.Title,
		Description:  m.Description,
		Completed:    m.Completed,
		CreatedAt:    m.CreatedAt,
		CompletedAt:  m.CompletedAt,
		AuthorUserID: m.AuthorUserID,
	}
}

package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_postgres_pool "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool"
)

func (r *TasksRepository) SaveTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.tasks (id, version, title, description, completed, created_at, completed_at, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;`

	m := domainToModel(task)
	row := r.pool.QueryRow(
		ctx, query,
		m.ID, m.Version, m.Title, m.Description, m.Completed, m.CreatedAt, m.CompletedAt, m.AuthorUserID,
	)

	if err := m.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf(
				"author user with id='%s': %w", task.AuthorUserID, core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	return modelToDomain(m), nil
}

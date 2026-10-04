package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_postgres_pool "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool"
)

func (r *TasksRepository) UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.tasks
	SET title=$1, description=$2, completed=$3, completed_at=$4, version=version+1
	WHERE id=$5 AND version=$6
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;`

	m := domainToModel(task)
	row := r.pool.QueryRow(ctx, query, m.Title, m.Description, m.Completed, m.CompletedAt, m.ID, m.Version)

	if err := m.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%s' concurrently accessed: %w", task.ID, core_errors.ErrConflict,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	return modelToDomain(m), nil
}

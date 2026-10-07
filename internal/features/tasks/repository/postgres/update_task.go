package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_postgres_pool "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool"
)

func (r *TasksRepository) UpdateTask(ctx context.Context, id uuid.UUID, input domain.TaskUpdate) (domain.Task, error) {
	query := `
	UPDATE todoapp.tasks
	SET title=$1, description=$2, completed=$3, completed_at=$4, version=version+1
	WHERE id=$5 AND version=$6
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;`

	row := r.pool.QueryRow(
		ctx, query,
		input.Title, input.Description, input.Completed, input.CompletedAt, id, input.Version,
	)

	var m TaskModel
	if err := m.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%s' concurrently accessed: %w: %w", id, err, core_errors.ErrConflict,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	return modelToDomain(m), nil
}

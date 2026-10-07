package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_postgres_pool "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool"
)

func (r *TasksRepository) CreateTask(ctx context.Context, input domain.TaskCreate) (domain.Task, error) {
	query := `
	INSERT INTO todoapp.tasks (title, description, author_user_id)
	VALUES ($1, $2, $3)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;`

	row := r.pool.QueryRow(ctx, query, input.Title, input.Description, input.AuthorUserID)

	var m TaskModel
	if err := m.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf(
				"author user with id='%s': %w: %w", input.AuthorUserID, err, core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	return modelToDomain(m), nil
}

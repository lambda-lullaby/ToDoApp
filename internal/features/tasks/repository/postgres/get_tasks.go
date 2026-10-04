package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	userID *uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;`

	args := []any{limit, offset}

	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id=$3")
		args = append(args, *userID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var m TaskModel
		if err := m.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		tasks = append(tasks, modelToDomain(m))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return tasks, nil
}

package tasks_postgres_repository

import (
	"context"
	"fmt"
	"strings"

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

	var query strings.Builder
	var args []any

	query.WriteString(`
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks`)
	if userID != nil {
		args = append(args, *userID)
		fmt.Fprintf(&query, " WHERE author_user_id=$%d", len(args))
	}
	query.WriteString(" ORDER BY id")
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

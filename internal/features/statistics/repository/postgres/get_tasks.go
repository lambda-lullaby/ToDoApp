package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

// GetTasks возвращает задачи, отфильтрованные по автору и полуоткрытому
// интервалу дат создания [from, to). Любой фильтр может быть nil.
func (r *StatisticsRepository) GetTasks(
	ctx context.Context,
	userID *uuid.UUID,
	from *time.Time,
	to *time.Time,
) ([]domain.Task, error) {
	var query strings.Builder
	args := []any{}
	conditions := []string{}

	query.WriteString(`
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks`)

	if userID != nil {
		args = append(args, *userID)
		conditions = append(conditions, fmt.Sprintf("author_user_id=$%d", len(args)))
	}
	if from != nil {
		args = append(args, *from)
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)))
	}
	if to != nil {
		args = append(args, *to)
		conditions = append(conditions, fmt.Sprintf("created_at<$%d", len(args)))
	}
	if len(conditions) > 0 {
		query.WriteString(" WHERE " + strings.Join(conditions, " AND "))
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

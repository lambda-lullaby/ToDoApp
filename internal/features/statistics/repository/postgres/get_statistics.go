package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

func (r *StatisticsRepository) GetStatistics(ctx context.Context, filter domain.StatisticsFilter) (domain.Statistics, error) {
	var query strings.Builder
	args := []any{}
	conditions := []string{}

	query.WriteString(`
	SELECT
		COUNT(*),
		COUNT(*) FILTER (WHERE completed),
		COUNT(*) FILTER (WHERE completed)::float8 / NULLIF(COUNT(*), 0) * 100,
		(EXTRACT(EPOCH FROM AVG(completed_at - created_at) FILTER (WHERE completed)) * 1000000)::bigint
	FROM todoapp.tasks`)

	if filter.UserID != nil {
		args = append(args, *filter.UserID)
		conditions = append(conditions, fmt.Sprintf("author_user_id=$%d", len(args)))
	}
	if filter.From != nil {
		args = append(args, *filter.From)
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)))
	}
	if filter.To != nil {
		args = append(args, *filter.To)
		conditions = append(conditions, fmt.Sprintf("created_at<$%d", len(args)))
	}
	if len(conditions) > 0 {
		query.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}
	query.WriteString(";")

	row := r.pool.QueryRow(ctx, query.String(), args...)

	var m StatisticsModel
	if err := m.Scan(row); err != nil {
		return domain.Statistics{}, fmt.Errorf("scan error: %w", err)
	}
	return modelToDomain(m), nil
}

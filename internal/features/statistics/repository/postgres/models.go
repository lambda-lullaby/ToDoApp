package statistics_postgres_repository

import (
	"time"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type StatisticsModel struct {
	TasksCreated                 int
	TasksCompleted               int
	TasksCompletedRate           *float64
	TasksAverageCompletionTimeUs *int64
}

type scanner interface {
	Scan(dest ...any) error
}

func (m *StatisticsModel) Scan(row scanner) error {
	return row.Scan(
		&m.TasksCreated,
		&m.TasksCompleted,
		&m.TasksCompletedRate,
		&m.TasksAverageCompletionTimeUs,
	)
}

func modelToDomain(m StatisticsModel) domain.Statistics {
	var avgTime *time.Duration
	if m.TasksAverageCompletionTimeUs != nil {
		d := time.Duration(*m.TasksAverageCompletionTimeUs) * time.Microsecond
		avgTime = &d
	}
	return domain.Statistics{
		TasksCreated:               m.TasksCreated,
		TasksCompleted:             m.TasksCompleted,
		TasksCompletedRate:         m.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}

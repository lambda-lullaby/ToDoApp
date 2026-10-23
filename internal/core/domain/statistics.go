package domain

import (
	"time"

	"github.com/google/uuid"
)

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

type StatisticsFilter struct {
	UserID *uuid.UUID
	From   *time.Time
	To     *time.Time `validate:"omitnil,statistics_to_after_from"`
}

func (f StatisticsFilter) Validate() error {
	return validate.Struct(f)
}

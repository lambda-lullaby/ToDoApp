package statistics_service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type StatisticsRepository interface {
	GetTasks(ctx context.Context, userID *uuid.UUID, from *time.Time, to *time.Time) ([]domain.Task, error)
}

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(statisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{statisticsRepository: statisticsRepository}
}

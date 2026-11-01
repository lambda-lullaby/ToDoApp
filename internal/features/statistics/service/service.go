package statistics_service

import (
	"context"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type StatisticsRepository interface {
	GetStatistics(ctx context.Context, filter domain.StatisticsFilter) (domain.Statistics, error)
}

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(statisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{statisticsRepository: statisticsRepository}
}

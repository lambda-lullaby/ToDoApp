package statistics_service

import (
	"context"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(ctx context.Context, filter domain.StatisticsFilter) (domain.Statistics, error) {
	if err := filter.Validate(); err != nil {
		return domain.Statistics{}, fmt.Errorf("validate statistics filter: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	statistics, err := s.statisticsRepository.GetStatistics(ctx, filter)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get statistics from repository: %w", err)
	}
	return statistics, nil
}

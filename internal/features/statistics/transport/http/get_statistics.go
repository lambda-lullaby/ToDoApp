package statistics_transport_http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type StatisticsResponse struct {
	TasksCreated               int      `json:"tasks_created"`
	TasksCompleted             int      `json:"tasks_completed"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(c *core_http.Context) error {
	userID, err := parseOptionalUUIDQuery(c, "user_id")
	if err != nil {
		return err
	}
	from, err := core_http.GetDateQueryParam(c.Request(), "from")
	if err != nil {
		return err
	}
	to, err := core_http.GetDateQueryParam(c.Request(), "to")
	if err != nil {
		return err
	}

	statistics, err := h.statisticsService.GetStatistics(c.Context(), domain.StatisticsFilter{
		UserID: userID,
		From:   from,
		To:     to,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, statisticsToResponse(statistics))
}

func statisticsToResponse(statistics domain.Statistics) StatisticsResponse {
	var avgTime *string
	if statistics.TasksAverageCompletionTime != nil {
		d := statistics.TasksAverageCompletionTime.String()
		avgTime = &d
	}
	return StatisticsResponse{
		TasksCreated:               statistics.TasksCreated,
		TasksCompleted:             statistics.TasksCompleted,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}

func parseOptionalUUIDQuery(c *core_http.Context, key string) (*uuid.UUID, error) {
	raw := c.QueryParam(key)
	if raw == "" {
		return nil, nil
	}
	value, err := uuid.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("parse `%s` query parameter: %v: %w", key, err, core_errors.ErrInvalidArgument)
	}
	return &value, nil
}

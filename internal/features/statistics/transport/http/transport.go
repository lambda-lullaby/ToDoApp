package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_http_server "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/server"
)

type StatisticsService interface {
	GetStatistics(ctx context.Context, userID *uuid.UUID, from *time.Time, to *time.Time) (domain.Statistics, error)
}

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{statisticsService: statisticsService}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{Method: http.MethodGet, Path: "/statistics", Handler: h.GetStatistics},
	}
}

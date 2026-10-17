package statistics_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

type fakeRepository struct {
	called bool
	tasks  []domain.Task
}

func (f *fakeRepository) GetTasks(context.Context, *uuid.UUID, *time.Time, *time.Time) ([]domain.Task, error) {
	f.called = true
	return f.tasks, nil
}

func TestGetStatistics_InvalidRange(t *testing.T) {
	from := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	cases := map[string]time.Time{
		"to before from": from.AddDate(0, 0, -1),
		"to equals from": from,
	}
	for name, to := range cases {
		t.Run(name, func(t *testing.T) {
			repo := &fakeRepository{}
			_, err := NewStatisticsService(repo).GetStatistics(context.Background(), nil, &from, &to)
			if !errors.Is(err, core_errors.ErrInvalidArgument) {
				t.Fatalf("expected ErrInvalidArgument, got %v", err)
			}
			if repo.called {
				t.Error("repository must not be called for invalid range")
			}
		})
	}
}

func TestGetStatistics_DelegatesToDomain(t *testing.T) {
	repo := &fakeRepository{tasks: []domain.Task{{CreatedAt: time.Now()}}}
	s, err := NewStatisticsService(repo).GetStatistics(context.Background(), nil, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.TasksCreated != 1 {
		t.Errorf("expected 1 task created, got %d", s.TasksCreated)
	}
}

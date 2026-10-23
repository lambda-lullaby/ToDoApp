package statistics_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

type fakeRepository struct {
	called bool
}

func (f *fakeRepository) GetStatistics(context.Context, domain.StatisticsFilter) (domain.Statistics, error) {
	f.called = true
	return domain.Statistics{}, nil
}

func TestGetStatistics_Validation(t *testing.T) {
	from := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	before := from.AddDate(0, 0, -1)
	after := from.AddDate(0, 1, 0)

	cases := []struct {
		name    string
		filter  domain.StatisticsFilter
		wantErr bool
	}{
		{"no filters", domain.StatisticsFilter{}, false},
		{"only from", domain.StatisticsFilter{From: &from}, false},
		{"only to", domain.StatisticsFilter{To: &from}, false},
		{"valid range", domain.StatisticsFilter{From: &from, To: &after}, false},
		{"to equals from", domain.StatisticsFilter{From: &from, To: &from}, true},
		{"to before from", domain.StatisticsFilter{From: &from, To: &before}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			repo := &fakeRepository{}
			_, err := NewStatisticsService(repo).GetStatistics(context.Background(), tc.filter)
			if tc.wantErr {
				if !errors.Is(err, core_errors.ErrInvalidArgument) {
					t.Fatalf("expected ErrInvalidArgument, got %v", err)
				}
				if repo.called {
					t.Error("repository must not be called for invalid filter")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !repo.called {
				t.Error("repository must be called for valid filter")
			}
		})
	}
}

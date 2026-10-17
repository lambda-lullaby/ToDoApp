package domain

import (
	"testing"
	"time"
)

func TestCreateStatistics_Empty(t *testing.T) {
	s := CreateStatistics(nil)

	if s.TasksCreated != 0 || s.TasksCompleted != 0 {
		t.Fatalf("expected zero counts, got %+v", s)
	}
	if s.TasksCompletedRate != nil {
		t.Errorf("expected nil rate, got %v", *s.TasksCompletedRate)
	}
	if s.TasksAverageCompletionTime != nil {
		t.Errorf("expected nil average time, got %v", *s.TasksAverageCompletionTime)
	}
}

func TestCreateStatistics_NoCompleted(t *testing.T) {
	now := time.Now()
	s := CreateStatistics([]Task{{CreatedAt: now}, {CreatedAt: now}})

	if s.TasksCreated != 2 || s.TasksCompleted != 0 {
		t.Fatalf("unexpected counts: %+v", s)
	}
	if s.TasksCompletedRate == nil || *s.TasksCompletedRate != 0 {
		t.Errorf("expected rate 0, got %v", s.TasksCompletedRate)
	}
	if s.TasksAverageCompletionTime != nil {
		t.Errorf("expected nil average time, got %v", *s.TasksAverageCompletionTime)
	}
}

func TestCreateStatistics_Mixed(t *testing.T) {
	created := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	done1 := created.Add(time.Minute)
	done2 := created.Add(2 * time.Minute)

	s := CreateStatistics([]Task{
		{CreatedAt: created, Completed: true, CompletedAt: &done1},
		{CreatedAt: created, Completed: true, CompletedAt: &done2},
		{CreatedAt: created},
		{CreatedAt: created},
	})

	if s.TasksCreated != 4 || s.TasksCompleted != 2 {
		t.Fatalf("unexpected counts: %+v", s)
	}
	if s.TasksCompletedRate == nil || *s.TasksCompletedRate != 50 {
		t.Errorf("expected rate 50, got %v", s.TasksCompletedRate)
	}
	if s.TasksAverageCompletionTime == nil || *s.TasksAverageCompletionTime != 90*time.Second {
		t.Errorf("expected average 1m30s, got %v", s.TasksAverageCompletionTime)
	}
}

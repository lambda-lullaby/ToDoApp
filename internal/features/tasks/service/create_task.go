package tasks_service

import (
	"context"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

func (s *TasksService) CreateTask(ctx context.Context, input domain.TaskCreate) (domain.Task, error) {
	if err := input.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task input: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	task, err := s.tasksRepository.CreateTask(ctx, input)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task in repository: %w", err)
	}
	return task, nil
}

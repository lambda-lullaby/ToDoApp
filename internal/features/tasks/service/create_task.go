package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	title string,
	description *string,
	authorUserID uuid.UUID,
) (domain.Task, error) {
	task := domain.CreateTask(title, description, authorUserID)

	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	savedTask, err := s.tasksRepository.SaveTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("save task in repository: %w", err)
	}
	return savedTask, nil
}

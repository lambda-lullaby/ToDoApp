package tasks_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type TasksRepository interface {
	SaveTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error)
	UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
	GetTasks(ctx context.Context, userID *uuid.UUID, limit *int, offset *int) ([]domain.Task, error)
}

type TasksService struct {
	tasksRepository TasksRepository
}

func NewTasksService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{tasksRepository: tasksRepository}
}

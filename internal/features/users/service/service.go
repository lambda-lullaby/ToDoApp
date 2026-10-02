package users_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
)

type UsersRepository interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
}

type UsersService struct {
	usersRepository UsersRepository
}

func NewUsersService(usersRepository UsersRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository}
}

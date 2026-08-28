package users_service

import (
	"context"
	"fmt"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

func (s *UsersService) CreateUser(ctx context.Context, fullName string, phoneNumber *string) (domain.User, error) {
	user := domain.CreateUser(fullName, phoneNumber)

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	savedUser, err := s.usersRepository.SaveUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("save user in repository: %w", err)
	}
	return savedUser, nil
}

package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

func (s *UsersService) PatchUser(ctx context.Context, id uuid.UUID, patch domain.UserPatch) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w: %w", err, core_errors.ErrInvalidArgument)
	}

	updatedUser, err := s.usersRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("update user in repository: %w", err)
	}
	return updatedUser, nil
}

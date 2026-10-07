package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID      uuid.UUID
	Version int64

	Title       string  `validate:"task_title"`
	Description *string `validate:"omitnil,task_description"`
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID uuid.UUID
}

func (t Task) Validate() error {
	if err := validate.Struct(t); err != nil {
		return err
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("`CompletedAt` can't be `nil` if `Completed`==`true`")
		}
		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("`CompletedAt` can't be before `CreatedAt`")
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("`CompletedAt` must be `nil` if `Completed`==`false`")
		}
	}

	return nil
}

type TaskCreate struct {
	Title        string    `validate:"task_title"`
	Description  *string   `validate:"omitnil,task_description"`
	AuthorUserID uuid.UUID `validate:"required"`
}

func (c TaskCreate) Validate() error {
	return validate.Struct(c)
}

type TaskUpdate struct {
	Version     int64
	Title       string
	Description *string
	Completed   bool
	CompletedAt *time.Time
}

func (t Task) ToUpdate() TaskUpdate {
	return TaskUpdate{
		Version:     t.Version,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CompletedAt: t.CompletedAt,
	}
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func (p TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("`Title` can't be null")
	}
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("`Completed` can't be null")
	}
	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", err)
	}

	tmp := *t
	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}
	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value
		switch {
		case !tmp.Completed:
			tmp.CompletedAt = nil
		case !t.Completed:
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}

	*t = tmp
	return nil
}

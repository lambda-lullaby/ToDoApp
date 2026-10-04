package domain

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

const (
	taskTitleMinLength       = 1
	taskTitleMaxLength       = 100
	taskDescriptionMinLength = 1
	taskDescriptionMaxLength = 1000
)

type Task struct {
	ID      uuid.UUID
	Version int64

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID uuid.UUID
}

func CreateTask(title string, description *string, authorUserID uuid.UUID) Task {
	return Task{
		ID:      uuid.New(),
		Version: 1,

		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,

		AuthorUserID: authorUserID,
	}
}

func (t Task) Validate() error {
	if length := utf8.RuneCountInString(t.Title); length < taskTitleMinLength || length > taskTitleMaxLength {
		return fmt.Errorf("`Title` must be between %d and %d characters long", taskTitleMinLength, taskTitleMaxLength)
	}
	if t.Description != nil {
		if length := utf8.RuneCountInString(*t.Description); length < taskDescriptionMinLength || length > taskDescriptionMaxLength {
			return fmt.Errorf(
				"`Description` must be between %d and %d characters long",
				taskDescriptionMinLength, taskDescriptionMaxLength,
			)
		}
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

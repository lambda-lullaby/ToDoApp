package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type TaskLimits struct {
	TitleMinLength       int
	TitleMaxLength       int
	DescriptionMinLength int
	DescriptionMaxLength int
}

func RegisterTaskValidation(limits TaskLimits) {
	validate.RegisterAlias(
		"task_title",
		fmt.Sprintf("min=%d,max=%d", limits.TitleMinLength, limits.TitleMaxLength),
	)
	validate.RegisterAlias(
		"task_description",
		fmt.Sprintf("min=%d,max=%d", limits.DescriptionMinLength, limits.DescriptionMaxLength),
	)
}

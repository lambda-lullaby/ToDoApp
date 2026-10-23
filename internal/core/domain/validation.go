package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = newValidator()

func newValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())
	_ = v.RegisterValidation("statistics_to_after_from", func(fl validator.FieldLevel) bool {
		filter, ok := fl.Parent().Interface().(StatisticsFilter)
		if !ok || filter.From == nil || filter.To == nil {
			return true
		}
		return filter.To.After(*filter.From)
	})
	return v
}

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

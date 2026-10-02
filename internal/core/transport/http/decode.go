package core_http

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"

	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func (c *Context) Bind(dest any) error {
	if err := json.NewDecoder(c.request.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if v, ok := dest.(validatable); ok {
		return wrapValidationErr(v.Validate())
	}
	return wrapValidationErr(requestValidator.Struct(dest))
}

func wrapValidationErr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("validate request: %v: %w", err, core_errors.ErrInvalidArgument)
}

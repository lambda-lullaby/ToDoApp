package core_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
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

package core_http

import (
	"fmt"
	"net/http"
	"time"

	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
)

const dateLayout = "2006-01-02"

func GetDateQueryParam(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}
	date, err := time.Parse(dateLayout, param)
	if err != nil {
		return nil, fmt.Errorf("parse `%s` query parameter: not a valid date: %v: %w", key, err, core_errors.ErrInvalidArgument)
	}
	return &date, nil
}

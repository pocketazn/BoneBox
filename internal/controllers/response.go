package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type HTTPError struct {
	Cause      error       `json:"-"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
}

type HTTPErrors struct {
	sync.Mutex // protects errors
	errors     []*HTTPError
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Message
	}

	return fmt.Sprintf("%s: %s", e.Message, e.Cause.Error())
}

func (e *HTTPErrors) add(statusCode int, displayMessage string, cause error) {
	e.append(&HTTPError{
		Cause:      cause,
		Message:    displayMessage,
		StatusCode: statusCode,
	})
}

func (e *HTTPErrors) append(err *HTTPError) {
	e.Lock()
	defer e.Unlock()

	e.errors = append(e.errors, err)
}

func (e *HTTPErrors) hasErrors() bool {
	e.Lock()
	defer e.Unlock()

	if e.errors == nil {
		return false
	}

	return len(e.errors) > 0
}

func respondError(ctx context.Context, w http.ResponseWriter, status int, message string, causer error) {
	resp := map[string]interface{}{
		"error": message,
	}

	if errors.Cause(causer) == sql.ErrNoRows {
		status = http.StatusNotFound
	}

	w.WriteHeader(status)

	// if we get spammed with too many 4xx errors, then revert to logging those on INFO
	log.Error(ctx, message, causer)

	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
}

func respondModel(ctx context.Context, w http.ResponseWriter, status int, model interface{}) {
	b, err := json.Marshal(model)
	if err != nil {
		respondError(ctx, w, 500, "error generating response", err)
	}

	w.WriteHeader(status)
	_, _ = w.Write(b)
}

//nolint:unparam
func respondModelWithError(ctx context.Context, w http.ResponseWriter, statusCode int, model map[string]interface{}, err *HTTPError) {
	httpErrors := HTTPErrors{}
	httpErrors.append(err)

	respondModelWithErrors(ctx, w, statusCode, model, &httpErrors)
}

//nolint:unparam
func respondModelWithErrors(ctx context.Context, w http.ResponseWriter, statusCode int, model map[string]interface{}, httpErrors *HTTPErrors) {
	maxStatusCode := statusCode

	// loop through errors, log as needed and find the highest status code
	if httpErrors != nil && httpErrors.hasErrors() {
		for _, err := range httpErrors.errors {
			if errors.Cause(err.Cause) == sql.ErrNoRows {
				err.StatusCode = http.StatusNotFound
			}

			if err.StatusCode > maxStatusCode {
				maxStatusCode = err.StatusCode
			}

			// if we get spammed with too many 4xx errors, then revert to logging those on INFO
			log.Error(ctx, err.Message, err.Cause)
		}

		model["errors"] = httpErrors.errors
	}

	jsonb, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"errors":[{"message":"error generating response"}]}`))

		return
	}

	w.WriteHeader(maxStatusCode)
	_, _ = w.Write(jsonb)
}
package usecase

import (
	"context"
	"net/http"

	"github.com/suzuito/server-go/internal/entity"
)

var keyContextLogEntry = "key_context_log_entry"

func SetContextLogEntry(r *http.Request, le *entity.LogEntry) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, keyContextLogEntry, le)
	return r.WithContext(ctx)
}

func GetContextLogEntry(r *http.Request) *entity.LogEntry {
	ctx := r.Context()
	le, _ := ctx.Value(keyContextLogEntry).(*entity.LogEntry)
	return le
}

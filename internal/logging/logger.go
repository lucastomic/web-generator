package logging

import (
	"context"
	"net/http"
)

type Logger interface {
	Request(context.Context, *http.Request, int)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

package types

import (
	"context"
	"net/http"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

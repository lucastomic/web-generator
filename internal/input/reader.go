package input

import (
	"net/http"

	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

type Reader interface {
	RetrieveInput(http.Request) (pagedata.PageData, error)
}

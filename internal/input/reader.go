package input

import (
	"net/http"

	"github.com/lucastomic/web-generator-service/internal/pagedata"
)

type Reader interface {
	RetrieveInput(http.Request) (pagedata.PageData, error)
}

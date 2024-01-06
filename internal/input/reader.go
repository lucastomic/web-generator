package input

import "github.com/lucastomic/web-generator/web-generator/internal/pagedata"

type Reader interface {
	RetrieveInput() pagedata.PageData
}

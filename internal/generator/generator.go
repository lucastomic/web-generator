package generator

import "github.com/lucastomic/web-generator/web-generator/internal/pagedata"

type Generator interface {
	Generate(pageData pagedata.PageData) error
}

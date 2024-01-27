package xmlinput

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/lucastomic/web-generator/web-generator/internal/logging"
	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

var inputFileParameter = "input"

type Reader struct {
	logging logging.Logger
}

func New(logging logging.Logger) Reader {
	return Reader{logging: logging}
}

func (reader Reader) RetrieveInput(req http.Request) (pagedata.PageData, error) {
	file, _, err := req.FormFile(inputFileParameter)
	if err != nil {
		return pagedata.PageData{}, err
	}
	defer file.Close()
	return reader.parseFile(file)
}

func (reader Reader) parseFile(file io.Reader) (pagedata.PageData, error) {
	var pageData pagedata.PageData
	decoder := xml.NewDecoder(file)
	decoder.Strict = false
	err := decoder.Decode(&pageData)
	return pageData, err
}

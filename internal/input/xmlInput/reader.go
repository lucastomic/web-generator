package xmlinput

import (
	"encoding/xml"
	"os"

	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

type Reader struct {
	filePath string
}

func New(filePath string) Reader {
	return Reader{filePath}
}

func (reader Reader) RetrieveInput() (pagedata.PageData, error) {
	file, err := os.Open(reader.filePath)
	if err != nil {
		return pagedata.PageData{}, err
	}
	defer file.Close()
	return reader.parseFile(file)
}

func (reader Reader) parseFile(file *os.File) (pagedata.PageData, error) {
	var pageData pagedata.PageData
	decoder := xml.NewDecoder(file)
	decoder.Strict = false
	err := decoder.Decode(&pageData)
	return pageData, err
}

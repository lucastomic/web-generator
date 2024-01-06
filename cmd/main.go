package main

import (
	"path/filepath"

	"github.com/lucastomic/web-generator/web-generator/internal/generator/templategenerator"
	xmlinput "github.com/lucastomic/web-generator/web-generator/internal/input/xmlInput"
)

func main() {
	inputPath, _ := filepath.Abs("./inputs/firstPage.xml")
	tmplPath, _ := filepath.Abs("./templates/template.html")
	reader := xmlinput.New(inputPath)
	generator := templategenerator.New(tmplPath)

	data, err := reader.RetrieveInput()
	if err != nil {
		panic(err)
	}
	err = generator.Generate(data)
	if err != nil {
		panic(err)
	}
}

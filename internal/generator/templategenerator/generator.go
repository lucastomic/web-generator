package templategenerator

import (
	"fmt"
	"html/template"
	"os"

	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

type Generator struct {
	templatePath string
}

func New(path string) Generator {
	return Generator{path}
}

func (gen Generator) Generate(pageData pagedata.PageData) error {
	tmpl, err := gen.createTemplate()
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("../generated/%s.html", pageData.Title)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, pageData)
}

func (gen Generator) createTemplate() (*template.Template, error) {
	funcMap := template.FuncMap{"mod": Mod, "sub": Sub}
	return template.New("template.html").
		Funcs(funcMap).
		ParseFiles(gen.templatePath)
}

package templategenerator

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/lucastomic/web-generator/web-generator/internal/logging"
	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

type Generator struct {
	templatePath string
	logging      logging.Logger
}

func New(templatePath string, logging logging.Logger) Generator {
	return Generator{templatePath, logging}
}

func (gen Generator) GenerateAndGetPaths(
	ctx context.Context,
	pageData pagedata.PageData,
) ([]string, error) {
	tmpl, err := gen.createTemplate()
	if err != nil {
		return []string{}, err
	}
	filePath := fmt.Sprintf("../tmp/%s.html", pageData.Title)
	file, err := os.Create(filePath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	err = tmpl.Execute(file, pageData)
	if err != nil {
		return []string{}, err
	}

	return []string{filePath}, nil
}

func (gen Generator) createTemplate() (*template.Template, error) {
	funcMap := template.FuncMap{"mod": Mod, "sub": Sub}
	return template.New("template.html").
		Funcs(funcMap).
		ParseFiles(gen.templatePath)
}

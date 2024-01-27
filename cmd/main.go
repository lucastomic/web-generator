package main

import (
	"path/filepath"

	"github.com/lucastomic/web-generator/web-generator/internal/generator/templategenerator"
	xmlinput "github.com/lucastomic/web-generator/web-generator/internal/input/xmlInput"
	"github.com/lucastomic/web-generator/web-generator/internal/logging"
	"github.com/lucastomic/web-generator/web-generator/internal/server"
	webprocessor "github.com/lucastomic/web-generator/web-generator/internal/webProcessor"
)

func main() {
	tmplPath, _ := filepath.Abs("../templates/template.html")
	logger := logging.NewLogrusLogger()
	reader := xmlinput.New(logger)
	generator := templategenerator.New(tmplPath, logger)
	webprocessor := webprocessor.New(logger, generator)
	server := server.New(":3001", webprocessor, reader, logger)
	server.Run()
}

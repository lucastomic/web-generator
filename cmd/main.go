package main

import (
	"path/filepath"

	"github.com/lucastomic/web-generator-service/internal/generator/templategenerator"
	xmlinput "github.com/lucastomic/web-generator-service/internal/input/xmlInput"
	"github.com/lucastomic/web-generator-service/internal/logging"
	"github.com/lucastomic/web-generator-service/internal/server"
	webprocessor "github.com/lucastomic/web-generator-service/internal/webProcessor"
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

package webprocessor

import (
	"context"
	"sync"

	"github.com/lucastomic/web-generator/web-generator/internal/generator"
	infraserviceconn "github.com/lucastomic/web-generator/web-generator/internal/infraServiceConn"
	"github.com/lucastomic/web-generator/web-generator/internal/logging"
	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

type WebProcessor interface {
	Process(context.Context, pagedata.PageData) error
}

type webProcessor struct {
	logging logging.Logger
	generator.Generator
}

func New(l logging.Logger, g generator.Generator) WebProcessor {
	return webProcessor{l, g}
}

func (wp webProcessor) Process(ctx context.Context, pageData pagedata.PageData) error {
	paths, err := wp.GenerateAndGetPaths(ctx, pageData)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(paths))

	for _, path := range paths {
		wg.Add(1)
		go sendToInfraService(path, errChan, &wg)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func sendToInfraService(path string, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	err := infraserviceconn.SendFileToInfraService(path)
	if err != nil {
		errChan <- err
	}
}

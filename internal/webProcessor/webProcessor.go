package webprocessor

import (
	"context"

	"github.com/lucastomic/web-generator-service/internal/generator"
	infraserviceconn "github.com/lucastomic/web-generator-service/internal/infraServiceConn"
	"github.com/lucastomic/web-generator-service/internal/logging"
	"github.com/lucastomic/web-generator-service/internal/pagedata"
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
	err = infraserviceconn.SendFilesToInfraService(paths)
	return err
}

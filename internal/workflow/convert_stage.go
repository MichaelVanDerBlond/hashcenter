package workflow

import (
	"context"
	"errors"

	"github.com/MichaelVanDerBlond/hashcenter/internal/convert"
)

type ConvertStage struct{}

func NewConvertStage() *ConvertStage {
	return &ConvertStage{}
}

func (c *ConvertStage) Name() string {
	return "convert"
}

func (c *ConvertStage) ShouldRun(wf *Context) bool {
	if wf == nil || wf.Analyze == nil {
		return false
	}

	return wf.Analyze.Conversion
}

func (c *ConvertStage) Run(ctx context.Context, wf *Context) error {
	_ = ctx

	if wf == nil {
		return errors.New("workflow context is nil")
	}

	result, err := convert.Run(wf.HashFile)
	if err != nil {
		return err
	}

	wf.Convert = result

	if result.Success && result.Output != "" {
		wf.HashFile = result.Output
	}

	return nil
}

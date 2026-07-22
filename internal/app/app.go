package app

import (
	"context"
	"errors"

	"github.com/MichaelVanDerBlond/hashcenter/internal/cli"
)

type App struct {
	ctx context.Context
}

func New(ctx context.Context) *App {
	if ctx == nil {
		ctx = context.Background()
	}

	return &App{
		ctx: ctx,
	}
}

func (a *App) Run() error {
	if a == nil {
		return errors.New("nil app")
	}

	return cli.Execute()
}

func Run() error {
	return New(context.Background()).Run()
}

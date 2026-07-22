package app

import (
	"context"

	"github.com/MichaelVanDerBlond/hashcenter/internal/cli"
)

type App struct {
	ctx context.Context
}

func New() *App {
	return &App{
		ctx: context.Background(),
	}
}

func (a *App) Run() error {
	return cli.Execute()
}

func Run() error {
	return New().Run()
}

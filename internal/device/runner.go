package device

import (
	"context"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func run(name string, args ...string) (string, error) {

	result, err := system.Exec(
		context.Background(),
		name,
		args...,
	)
	if err != nil {
		return "", err
	}

	return result.Stdout, nil
}

package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/version"
)

func Execute() error {
	fmt.Printf("%s %s\n", version.Name, version.Version)
	return nil
}

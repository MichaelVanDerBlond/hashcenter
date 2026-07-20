package cli

import "github.com/MichaelVanDerBlond/hashcenter/internal/web"

func Serve() error {
	return web.Run(":1111")
}

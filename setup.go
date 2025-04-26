package valkey

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/plugin"
)

func init() {
	caddy.RegisterPlugin("valkey", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	r, err := valkeyParse(c)
	if err != nil {
		return plugin.Error("valkey", err)
	}
}

// func valkeyParse(c *caddy.Controller) (*)

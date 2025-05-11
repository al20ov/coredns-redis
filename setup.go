package valkey

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	caddy.RegisterPlugin("valkey", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	v, err := valkeyParse(c)
	if err != nil {
		return plugin.Error("valkey", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		v.Next = next
		return v
	})

	return nil
}

func valkeyParse(c *caddy.Controller) (*Valkey, error) {
	v := Valkey{
		keyPrefix: "",
		keySuffix: "",
		Ttl:       300,
	}

	var err error

	for c.Next() {
		if c.NextBlock() {
			for {
				switch c.Val() {
				case "address":
					if !c.NextArg() {
						return &Valkey{}, c.ArgErr()
					}
				}
			}
		}
	}
}

package kafka

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("kafka", setup) }

func setup(c *caddy.Controller) error {
	k, err := paramsParse(c)

	if err != nil {
		return plugin.Error("kafka err", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		k.Next = next
		return k
	})

	return nil

}

func paramsParse(c *caddy.Controller) (Kafka, error) {
	kafka := Kafka{}
	for c.Next() {
		args := c.RemainingArgs()

		switch len(args) {
		case 0:
		case 1:
			kafka.hosts = args[0]
		default:
			return kafka, c.ArgErr()
		}

		for c.NextBlock() {
			switch c.Val() {
			case "username":
				args := c.RemainingArgs()
				if len(args) != 1 {
					return kafka, c.ArgErr()
				}
				kafka.username = args[0]
			case "password":
				args := c.RemainingArgs()
				if len(args) != 1 {
					return kafka, c.ArgErr()
				}
				kafka.password = args[0]
			default:
				return kafka, c.ArgErr()
			}
		}
	}
	return kafka, nil
}

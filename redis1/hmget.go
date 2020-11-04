package redis1

import (
	"examples/config"
	"examples/utils"
	"github.com/micro/cli"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

var (
	prometheusAddr string
)

var loadFlags = micro.Action(func(c *cli.Context) error {
	prometheusAddr = c.String("prometheus_address")
	return nil
})

func main() {
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{Name: "prometheus_address", Usage: "The prometheus service"},
		),
	)
	service.Init(loadFlags)

	utils.Micro.Init(service)
	utils.Micro.LoadSource()
	utils.Prometheus(prometheusAddr)

	
}

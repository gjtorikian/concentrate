package concentrate

import (
	"github.com/codegangsta/cli"

	"fmt"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "concentrate"
	app.Usage = "Disables WiFi for a given amount of time"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "time, t",
			Value: "",
			Usage: "The length of time to disable WiFi",
		},
	}
	app.Action = func(c *cli.Context) {
		run(c)
	}

	return app
}

type runFunc func(*cli.Context)

var run = func(c *cli.Context) {
  fmt.Println("Hello, World")
}

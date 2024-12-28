package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Webhealthchecker",
		Usage: "Checks if a website is healthy",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port name to check",
				Required: false,
			},
		},

		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			if ctx.String("port") == "" {
				port = "80"
			}

			status := Check(ctx.String("domain"), port)
			fmt.Println(status)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
}

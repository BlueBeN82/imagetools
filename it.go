package main

import (
	"context"
	"os"

	"github.com/blueben82/imagetools/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:    "imagetools",
		Version: "0.0.5",
		Usage:   "A tiny helper library, written in golang, that serves some utilities that can be used to build more robust docker containers",

		Commands: []*cli.Command{
			{
				Name:  "logs",
				Usage: "Super easy logging",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "debug",
					},
					&cli.BoolFlag{
						Name: "info",
					},
				},
				Action: commands.Logs,
			},
			{
				Name:  "requires",
				Usage: "Validates environment variables and terminates script if they dont match your expectations. By default it expects the given env variable to be a not empty string.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "validate",
						Usage: "Select one of int, float, bool, string or regex as validator",
					},
					&cli.StringFlag{
						Name:  "min",
						Usage: "The minimum accepted value for int or float",
					},
					&cli.StringFlag{
						Name:  "max",
						Usage: "The maximum accepted value for int or float",
					},
					&cli.StringFlag{
						Name:  "pattern",
						Usage: "The regex pattern used to validate the value",
					},
					&cli.StringFlag{
						Name:  "example",
						Usage: "An example value to be used in error message",
					},
				},
				Action: commands.Requires,
			},
			{
				Name:   "waits-for",
				Usage:  "Waits for an external service to become available.",
				Action: commands.WaitsFor,
			},
		},
	}

	app.Run(context.Background(), os.Args)
}

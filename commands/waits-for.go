package commands

import (
	"github.com/urfave/cli"
	"fmt"
)

func WaitsFor(c *cli.Context) error {
	fmt.Fprintf(c.App.Writer, ":wave: over here, eh\n")
	return nil
}

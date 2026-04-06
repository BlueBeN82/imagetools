package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func WaitsFor(ctx context.Context, cmd *cli.Command) error {
	fmt.Fprintf(cmd.Root().Writer, ":wave: over here, eh\n")
	return nil
}

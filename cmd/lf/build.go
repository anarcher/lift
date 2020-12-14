package main

import (
	"fmt"
	"os"

	"lift/pkg/app"

	"github.com/spf13/cobra"
)

func buildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "build",
		Short:   "Print configuration",
		Aliases: []string{"b"},
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		err := app.Build(app.NewBuildOpts())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return cmd
}

package main

import (
	"fmt"
	"lift/pkg/app"
	"os"

	"github.com/spf13/cobra"
)

func checkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check",
		Short:   "Check configuration",
		Aliases: []string{"c"},
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		err := app.Check(app.NewCheckOpts())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "OK")
	}

	return cmd
}

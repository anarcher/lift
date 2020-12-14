package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"lift/pkg/app"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes lift directory in the current directory.",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		err := app.Init(app.NewInitOpts())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return cmd
}

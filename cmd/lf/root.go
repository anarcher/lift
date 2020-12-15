package main

import "github.com/spf13/cobra"

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lf",
		Short: "lift: a stupid CLI",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Help()
	}

	cmd.AddCommand(initCmd())
	cmd.AddCommand(buildCmd())
	cmd.AddCommand(checkCmd())
	cmd.AddCommand(pkgCmd())

	return cmd
}

package main

import (
	"fmt"
	"lift/pkg/cueutil"
	"os"

	"github.com/anarcher/cue-bundler/pkg/cb"
	"github.com/spf13/cobra"
)

func pkgCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pkg",
		Short: "Install new packages",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Help()
	}
	cmd.AddCommand(pkgInstallCmd())
	cmd.AddCommand(pkgUpdateCmd())

	return cmd
}

func pkgInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install new packages",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		modDir := cueutil.FindModDirPath()
		uris := args
		if err := cb.Install(modDir, uris); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return cmd
}

func pkgUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update all or specific packages",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		modDir := cueutil.FindModDirPath()
		uris := args
		if err := cb.Update(modDir, uris); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return cmd
}

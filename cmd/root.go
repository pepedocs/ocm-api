package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ocm-api cli",
	Short: "CLI for inspecting OCM APIs",
	Long:  "CLI for inspecting OCM APIs",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

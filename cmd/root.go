package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ocm-api cli",
	Short: "CLI for querying OCM API information",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

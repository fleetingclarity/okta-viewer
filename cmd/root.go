package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "okta-viewer",
	Short: "CLI for Okta user and group information",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiger",
	Short: "Tiger is a custom Reviewpad CLI",
	Long:  "Tiger is a custom Reviewpad CLI that will automate a lot of dev tasks.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

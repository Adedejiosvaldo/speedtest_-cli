package cmd

import (
	"fmt"
	"os"

	"github.com/adedejiosvaldo/terminal_speedtest/helpers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Speedup",
	Short: "Speedup is a cli application for performing internet speed test",
	Long:  "Speedup is a cli application for performing internet speed test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("running zero")

		helpers.PingServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}

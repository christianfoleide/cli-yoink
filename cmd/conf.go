package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(confCmd)
}

var (
	confCmd = &cobra.Command{
		Use: "config",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("Config command called")
		},
	}
)

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	confCmd.AddCommand(listCmd)
}

var (
	listCmd = &cobra.Command{
		Use: "list",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("List command called")
		},
	}
)

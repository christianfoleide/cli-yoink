package cmd

import (
	"fmt"

	"github.com/christianfoleide/yoink/yoink"
	"github.com/spf13/cobra"
)

func init() {
	confCmd.AddCommand(listCmd)
}

var (
	listCmd = &cobra.Command{
		Use: "list",
		Run: func(c *cobra.Command, args []string) {

			f := yoink.NewFilehandler("config.json")
			if err := f.ListConfig(); err != nil {
				fmt.Printf("error: %s", err)
			}

		},
	}
)

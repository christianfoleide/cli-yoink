package cmd

import (
	"fmt"
	"github.com/christianfoleide/yoink/util"
	"os"
	"github.com/christianfoleide/yoink/client"
	"github.com/spf13/cobra"
)

var (
	aliases = []string{"yoink"}

	Pretty bool
	Method bool

	rootCmd = &cobra.Command{
		Use:     "yo",
		Aliases: aliases,
		Short:   "A network CLI for get, post and put requests to an API",
		Long:    "A network CLI for GET, POST, and PUT requests to an API where you can specify json files to send as test data. Defaults to GET",
		Args: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			c := client.NewClient(args[0], args[1])
			b, err := c.DoRequest()
			if err != nil {
				fmt.Println(err)
			}

			util.PrettyPrint(b)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Pretty, "pretty", "p", false, "Pretty print request result")
	rootCmd.PersistentFlags().BoolVarP(&Method, "method", "m", false, "Specifies another request method")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

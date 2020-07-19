package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var (

	aliases = []string{"yoink"}

	Pretty bool
	Method bool

	rootCmd = &cobra.Command{
		Use: "yo",
		Aliases: aliases,
		Short: "A network CLI for get, post and put requests to an API",
		Long: "A network CLI for GET, POST, and PUT requests to an API where you can specify json files to send as test data. Defaults to GET",
		Run: func(cmd *cobra.Command, args []string) {



		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Pretty, "pretty", "p", false, "Pretty print request result")
	rootCmd.PersistentFlags().BoolVarP(&Method, "method", "m", false, "Specify another request method passed as first argument")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strings"
	"github.com/christianfoleide/yoink/methods"
	"github.com/christianfoleide/yoink/util"
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
		Args: func(cmd *cobra.Command, args []string) error {

			if !Method { //first and only argument is URL
				return util.ValidateUrl(args[0])
			}
			//first argument is request-method, second argument is URL
			return util.ValidateArgs(args)
		},
		Run: func(cmd *cobra.Command, args []string) {

			if !Method { //default get
				dest := args[0]
				result, err := methods.Get(dest)

				if err != nil {
					fmt.Println(err)
					return
				}

				if Pretty {
					util.PrettyPrint(result)
				} else {
					fmt.Println(string(result))
				}
				return
			}
			//args are already validated

			method := strings.ToUpper(args[0])
			dest := args[1]
			file := args[2]

			switch method {
			case http.MethodPost:
				result, err := methods.Post(dest, file)
				if err != nil {
					fmt.Println(err)
					return
				}
				if Pretty {
					util.PrettyPrint(result)
					return
				}
				fmt.Println(string(result))
				return
				
			case http.MethodPut:
				result, err := methods.Put(dest, file)
				if err != nil {
					fmt.Println(err)
					return
				}
				if Pretty {
					util.PrettyPrint(result)
					return
				}
				fmt.Println(string(result))
				return
			}
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

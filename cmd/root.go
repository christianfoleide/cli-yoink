package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/christianfoleide/yoink/validation"
	"github.com/christianfoleide/yoink/yoink"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var (
	changeMethod bool
	useConfig    bool
	configFile   = "config.json"

	deadline = time.Second * 4

	rootCmd = &cobra.Command{
		Use:   "yoink",
		Short: "A network CLI tool for sending requests to a resource",
		Args:  validate,
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {

	f := yoink.NewFilehandler(configFile)

	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()

	if useConfig {
		c, err := f.ConfigToClient()

		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}

		b, err := c.DoRequest(ctx)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}

		PrettyPrint(b)
		return
	}

	if changeMethod {

		method := strings.ToUpper(args[0])
		uri := args[1]

		payloadFile := args[2]
		payload, err := yoink.ReadFile(payloadFile)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		c := yoink.NewClient(method, uri, payload)
		b, err := c.DoRequest(ctx)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		PrettyPrint(b)

		return
	}
	//default GET
	c := yoink.DefaultClient(args[0])
	b, err := c.DoRequest(ctx)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	PrettyPrint(b)
	return
}

func validate(cmd *cobra.Command, args []string) error {

	if useConfig {
		return nil
	}

	if changeMethod {
		if err := validation.ValidateNonDefault(args); err != nil {
			return err
		}

	} else {
		if err := validation.ValidateDefault(args); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.Flags().BoolVar(&useConfig, "use-config", false, "do request specified in configuration file")
	rootCmd.Flags().BoolVarP(&changeMethod, "method", "m", false, "specifies another request method")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PrettyPrint(response []byte) {
	pretty := pretty.Pretty(response)
	fmt.Println(string(pretty))
}

package cmd

import (
	"fmt"

	"github.com/christianfoleide/yoink/yoink"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(confCmd)

	confCmd.Flags().StringVar(&host, "set-hostname", "", "set hostname in configuration")
	confCmd.Flags().StringVar(&port, "set-port", "", "set port in configuration")
	confCmd.Flags().StringVar(&path, "set-path", "", "set URI path in configuration")
	confCmd.Flags().StringVar(&method, "set-method", "", "set request method in configuration")
	confCmd.Flags().StringVar(&payload, "set-payload", "", "set filename for payload in configuration")

}

var (
	host    string
	port    string
	path    string
	method  string
	payload string

	confCmd = &cobra.Command{
		Use: "config",
		Run: func(c *cobra.Command, args []string) {

			changes := determineChanges()
			if len(changes) > 0 { //flags used

				f := yoink.NewFilehandler("config.json")
				f.WriteChanges(changes)

			} else {
				fmt.Println("no flags present, use yoink config -h")
			}

		},
	}
)

func determineChanges() map[string]interface{} {

	changes := make(map[string]interface{})

	if len(host) > 0 {
		changes["hostname"] = host
	}

	if len(port) > 0 {
		changes["port"] = port
	}

	if len(path) > 0 {
		changes["path"] = path
	}

	if len(method) > 0 {
		changes["method"] = method
	}

	if len(payload) > 0 {
		changes["payload"] = payload
	}

	return changes

}

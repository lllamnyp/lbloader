/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/lllamnyp/lbloader/pkg/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Config{URL: args[0]})
		c.Call()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")
	clientCmd.PersistentFlags().String("listen-addr", "0.0.0.0", "The address to which to bind.")
	clientCmd.PersistentFlags().Int64P("listen-port", "p", 8443, "The port on which to listen.")
	clientCmd.PersistentFlags().String("cert", "/etc/ssl/promidc/tls.crt", "The path to the certificate of the TLS server.")
	clientCmd.PersistentFlags().String("key", "/etc/ssl/promidc/tls.key", "The path to the private key of the TLS server.")
	clientCmd.PersistentFlags().String("oidc-endpoint", "https://localhost", "The OIDC endpoint for token inspection.")
	clientCmd.PersistentFlags().String("datasource", "prometheus", "The type of the datasource. Currently 'prometheus' and 'loki' are supported.")
	clientCmd.PersistentFlags().String("datasource-url", "http://localhost:9090", "The URL of the datasource server.")
	clientCmd.PersistentFlags().Bool("serve-insecure", false, "This option disables TLS on the server.")
	clientCmd.PersistentFlags().Bool("passthrough", false, "Authorize all requests. Use only for debugging.")
	clientCmd.PersistentFlags().BoolP("insecure-skip-verify", "k", false, "This option skips certificate verification on prometheus and the authorization server.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

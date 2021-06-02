/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/kostis-codefresh/headmast/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server component",
	Long: `Starts the server component that serves both the UI and the API at /api
	You can specify the host and port with the respective flags`,
	Example: "headmast server --address=127.0.0.1 --port=9000",
	Run: func(cmd *cobra.Command, args []string) {
		a := server.App{}

		a.Run(address + ":" + port)
	},
}

var (
	address string
	port    string
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "Address to listen")
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen")
}

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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kostis-codefresh/headmast/model"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var appCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		app := model.Application{
			Name:    name,
			Version: version,
		}

		jsonValue, _ := json.Marshal(app)

		url := "http://" + address + ":" + port + "/api/apps"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Println("Non-OK HTTP status:", resp.StatusCode)
			return
		}

		log.Println("Response status of api.github.com:", resp.Status)

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			log.Println(err.Error())
			return
		}

		returnedApp := model.Application{}
		err = json.Unmarshal(buf.Bytes(), &returnedApp)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(returnedApp)
	},
}

var (
	name    string
	version string
)

func init() {
	applicationCmd.AddCommand(appCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	appCreateCmd.Flags().StringVarP(&name, "name", "n", "example-app", "Name of the application")
	appCreateCmd.Flags().StringVarP(&version, "version", "b", "0.1", "Version of the application")

}

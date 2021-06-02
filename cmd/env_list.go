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
var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all applications",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		url := "http://" + address + ":" + port + "/api/envs"
		resp, err := http.Get(url)
		log.Println("Calling " + url)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("Called")
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Non-OK HTTP status:", resp.StatusCode)
			// return assetDetails, errors.New("Could not access " + url)
			return
		}

		log.Println("Response status of api.github.com:", resp.Status)

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return
		}

		var envs []model.Environment
		err = json.Unmarshal(buf.Bytes(), &envs)

		if err != nil {
			return
		}
		log.Printf("Unmarshaled: %v", envs)
	},
}

func init() {
	envCmd.AddCommand(envListCmd)

}

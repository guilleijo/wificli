/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// connCmd represents the conn command
var connCmd = &cobra.Command{
	Use:   "conn",
	Short: "Connect to a wifi network",
	Long:  `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Name and password required")
			return
		}

		fmt.Println("Connecting to wifi...")

		name := args[0]
		password := args[1]
		cmdStr := fmt.Sprintf("networksetup -setairportnetwork en0 %s %s", name, password)

		turnOffCmd := exec.Command("bash", "-c", cmdStr)
		out, err := turnOffCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(connCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connCmd.Flags().StringP("name", "n", "", "SSID of network")
	// connCmd.Flags().StringP("password", "p", "", "Password of network")
}

package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn wifi on",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Turning wifi on")

		turnOnCmd := exec.Command("bash", "-c", "networksetup -setairportpower en0 on")
		out, err := turnOnCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}

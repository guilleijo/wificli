package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn wifi off",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Turning wifi off")

		turnOffCmd := exec.Command("bash", "-c", "networksetup -setairportpower en0 off")
		out, err := turnOffCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}

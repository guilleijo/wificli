package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/guilleijo/wificli/utils"
	"github.com/spf13/cobra"
)

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn wifi off",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Prefix = "Turning wifi off "
		s.FinalMSG = "Wifi is off!\n"
		s.Start()
		defer s.Stop()

		wifiPort := utils.GetWifiPort()
		cmdStr := fmt.Sprintf("networksetup -setairportpower %s off", wifiPort)
		turnOffCmd := exec.Command("bash", "-c", cmdStr)
		out, err := turnOffCmd.Output()
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		if string(out) != "" {
			fmt.Printf("%s", out)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}

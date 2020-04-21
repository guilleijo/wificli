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

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn wifi on",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Prefix = "Turning wifi on "
		s.FinalMSG = "Wifi is on!\n"
		s.Start()
		defer s.Stop()

		wifiPort := utils.GetWifiPort()
		cmdStr := fmt.Sprintf("networksetup -setairportpower %s on", wifiPort)
		turnOnCmd := exec.Command("bash", "-c", cmdStr)
		out, err := turnOnCmd.Output()
		utils.HandleError(err)

		if string(out) != "" {
			fmt.Printf("%s", out)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}

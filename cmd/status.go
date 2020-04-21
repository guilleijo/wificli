package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/guilleijo/wificli/utils"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Wifi status",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Prefix = "Getting wifi status: "
		s.Start()

		wifiStatusCmd := exec.Command("bash", "-c", "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/A/Resources/airport -I")
		out, err := wifiStatusCmd.Output()
		utils.HandleError(err)

		s.Stop()
		fmt.Printf("Wifi status:\n %s\n", out)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

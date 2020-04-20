package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
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

		turnOnCmd := exec.Command("bash", "-c", "networksetup -setairportpower en0 on")
		out, err := turnOnCmd.Output()
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
	rootCmd.AddCommand(onCmd)
}

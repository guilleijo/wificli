package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
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

		turnOffCmd := exec.Command("bash", "-c", "networksetup -setairportpower en0 off")
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

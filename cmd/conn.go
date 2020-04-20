package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"wificli/utils"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var connCmd = &cobra.Command{
	Use:   "conn",
	Short: "Connect to a wifi network",
	Long:  `Select a wifi network to connect.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")

		if name == "" {
			name, err = selectWifi()
		}

		if password == "" {
			password, err = utils.Prompt("Enter your password", true)
		}

		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Prefix = fmt.Sprintf("Connecting to %s: ", name)
		s.Start()

		cmdStr := fmt.Sprintf("networksetup -setairportnetwork en0 %s %s", name, password)
		connectCmd := exec.Command("bash", "-c", cmdStr)
		out, err := connectCmd.Output()
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		s.Stop()
		outStr := string(out)
		if outStr != "" {
			errMsg := strings.Split(outStr, "\n")[0]
			fmt.Printf("%s\n", errMsg)
			os.Exit(1)
		}

		fmt.Printf("Success! Connected to %s\n", name)
	},
}

func selectWifi() (string, error) {
	wifiList := utils.ListWifiNetworks()

	prompt := promptui.Select{
		Label: "Select wifi network",
		Items: wifiList,
		Size:  len(wifiList),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return result, nil
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(connCmd)

	// Local flags
	connCmd.Flags().StringP("name", "n", "", "SSID of network")
	connCmd.Flags().StringP("password", "p", "", "Password of network")
}

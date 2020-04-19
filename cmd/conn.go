package cmd

import (
	"fmt"
	"os/exec"
	"wificli/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var connCmd = &cobra.Command{
	Use:   "conn",
	Short: "Connect to a wifi network",
	Long:  `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")

		if name == "" {
			name, err = selectWifi()
		}

		if password == "" {
			password, err = utils.Prompt("Enter your password")
		}

		fmt.Printf("Connecting to %s...\n", name)
		cmdStr := fmt.Sprintf("networksetup -setairportnetwork en0 %s %s", name, password)

		connectCmd := exec.Command("bash", "-c", cmdStr)
		_, err = connectCmd.Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("Connected to %s\n", name)
	},
}

func selectWifi() (string, error) {
	wifiList := utils.ListWifiNetworks()[1:] // Remove header

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

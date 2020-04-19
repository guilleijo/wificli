package cmd

import (
	"fmt"
	"strings"
	"wificli/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available wifi networks",
	Run: func(cmd *cobra.Command, args []string) {
		wifiList := utils.ListWifiNetworks()
		stringSlices := strings.Join(wifiList[:], "\n")
		fmt.Printf("%s\n", stringSlices)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

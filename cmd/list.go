package cmd

import (
	"fmt"
	"strings"

	"github.com/guilleijo/wificli/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available wifi networks",
	Run: func(cmd *cobra.Command, args []string) {
		wifiList := utils.ListWifiNetworks()
		stringSlices := strings.Join(wifiList[:], "\n")
		fmt.Println(stringSlices)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ddo/go-fast"
	"github.com/spf13/cobra"
)

var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "Download speed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Download speed: ")
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Start()
		defer s.Stop()
		testFast()
	},
}

func init() {
	rootCmd.AddCommand(speedCmd)
}

func testFast() {
	fastCom := fast.New()

	err := fastCom.Init()
	if err != nil {
		panic(err)
	}

	urls, err := fastCom.GetUrls()
	if err != nil {
		panic(err)
	}

	KbpsChan := make(chan float64)
	go func() {
		for Kbps := range KbpsChan {
			fmt.Printf("%.2f Mbps\n", Kbps/1000)
		}
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		panic(err)
	}
}

package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
)

// ListWifiNetworks returns list of wifi network names
func ListWifiNetworks() []string {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "Searching networks: "
	s.Start()
	defer s.Stop()

	listWifiCmd := exec.Command("bash", "-c", "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/A/Resources/airport scan")
	wifiList, err := listWifiCmd.Output()

	var list []string
	wifiListArray := strings.Split(string(wifiList), "\n")[1:]
	for _, network := range wifiListArray {
		networkFields := strings.Fields(network)
		if len(networkFields) > 0 {
			list = append(list, networkFields[0])
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return list
}

// Prompt shows `msg` and accepts input
func Prompt(msg string, isPassword bool) (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Value can't be empty")
		}
		return nil
	}

	var prompt promptui.Prompt
	if isPassword == true {
		prompt = promptui.Prompt{
			Label:    msg,
			Validate: validate,
			Mask:     '*',
		}
	} else {
		prompt = promptui.Prompt{
			Label:    msg,
			Validate: validate,
		}
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

package main

import (
	// "fmt"
	"os/exec"
	"strings"
)

func scanWiFis() error {
	cmd := exec.Command("iwctl", "station", "wlan0", "scan")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func fetchWiFis() ([]string, error) {
	cmd := exec.Command("iwctl", "station", "wlan0", "get-networks")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(output), "\n"), nil
}

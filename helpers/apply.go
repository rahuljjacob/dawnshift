package helpers

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func ApplyFilter(config *Config) {
	cmd := exec.Command("pgrep", "-x", "hyprsunset")
	err := cmd.Run()
	if err != nil {
		startCmd := exec.Command("hyprsunset")
		err := startCmd.Start()
		if err != nil {
			fmt.Println("Error starting hyprsunset:", err)
			return
		}
	}

	curTime := time.Now()
	tempComparisonTime := curTime.Add(2 * time.Minute)
	comparisonTime := normalizeTime(tempComparisonTime)
	tempApplied := false

	for _, periodTimes := range config.Periods {
		// If period goes through the 12:00
		startTime := normalizeTime(periodTimes.StartTime)
		endTime := normalizeTime(periodTimes.EndTime)
		if endTime.Before(startTime) {
			if comparisonTime.After(startTime) || comparisonTime.Before(endTime) {
				cmd := exec.Command(
					"hyprctl",
					"hyprsunset",
					"temperature",
					strconv.Itoa(periodTimes.Temperature),
				)
				_, err := cmd.Output()
				if err != nil {
					fmt.Println(err)
				}
				tempApplied = true
				break
			}
		} else {
			if comparisonTime.After(startTime) && comparisonTime.Before(endTime) {
				cmd := exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(periodTimes.Temperature))
				_, err := cmd.Output()
				if err != nil {
					fmt.Println(err)
				}
				tempApplied = true
				break
			}
		}

		if !tempApplied {
			cmd := exec.Command(
				"hyprctl",
				"hyprsunset",
				"temperature",
				strconv.Itoa(config.Default.Temperature),
			)
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			tempApplied = true
		}
	}
}

func normalizeTime(t time.Time) time.Time {
	return time.Date(0, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
}

package helpers

import (
	"fmt"
	"os/exec"
	"time"
	"strconv"
)

func ApplyFilter(config *Config){
	curTime := time.Now()
	comparisonTime := curTime.Add(2 * time.Minute)
	tempApplied := false

	for _, periodTimes := range config.Periods {
		// If period goes through the 12:00
		if periodTimes.EndTime.Before(periodTimes.StartTime){ 
			if comparisonTime.After(periodTimes.StartTime) || comparisonTime.Before(periodTimes.EndTime) {
				cmd := exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(periodTimes.Temperature))
				stdout, err := cmd.Output()
				if err != nil {
					fmt.Println(err)
				}
				tempApplied = true
				fmt.Println(string(stdout))
				break
			}
		} else {
			if comparisonTime.After(periodTimes.StartTime) && comparisonTime.Before(periodTimes.EndTime) {
				cmd := exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(periodTimes.Temperature))
				stdout, err := cmd.Output()
				if err != nil {
					fmt.Println(err)
				}
				tempApplied = true
				fmt.Println(string(stdout))
				break
			}
		}

		if tempApplied == false {
			cmd := exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(config.Default.Temperature))
			_ , err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			tempApplied = true
		}
	}
}
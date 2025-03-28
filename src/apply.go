package src

import (
	"os/exec"
	"time"
	"strconv"
)

func ApplyFilter(config *Config){
	curTime := time.Now()
	comparisonTime := curTime.Add(2 * time.Minute)

	for _, periodTimes := range config.Periods {
		if comparisonTime.After(periodTimes.StartTime) && comparisonTime.Before(periodTimes.EndTime) {
			cmd := exec.Command("hyprctl hyprsunset temperature " + strconv.Itoa(periodTimes.Temperature))
			err := cmd.Run()
			if err != nil {
				println("Error applying filter:", err.Error())
			}
			break
		}
	}
}
package src

import (
	"fmt"
	"os"
)



func WriteTimer(config *Config){
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	file, err := os.Create(home + "/.config/systemd/user/sunset.timer")
	if err != nil {
		fmt.Println("Error writing to timer file")
	}

	uniqueTimes := make(map[string]struct{})

	timeFormat := "15:04:05"
	file.WriteString("[Unit]\nDescription=Apply screen filter on schedule\n[Timer]\n")
	for _, signalTime := range config.Periods {
		startTime := "OnCalendar=*-*-* "+signalTime.StartTime.Format(timeFormat)+"\n"
		endTime := "OnCalendar=*-*-* "+signalTime.EndTime.Format(timeFormat)+"\n"

		if _, exists := uniqueTimes[startTime]; !exists {
			file.WriteString(startTime)
			uniqueTimes[startTime] = struct{}{}
		}
		if _, exists := uniqueTimes[endTime]; !exists {
			file.WriteString(endTime)
			uniqueTimes[endTime] = struct{}{}
		}
	}
}

package helpers

import (
	"fmt"
	"os"
)

func writeTimer(config *Config){
	file, err := os.Create("/etc/systemd/system/sunset.timer")
	if err != nil {
		fmt.Println("Error writing to timer file")
	}

	uniqueTimes := make(map[string]bool)

	timeFormat := "15:04:05"
	file.WriteString("[Unit]\nDescription=Apply screen filter on schedule\n\n[Timer]\n")
	for _, periodTimes := range config.Periods {
		startTime := "OnCalendar=*-*-* "+periodTimes.StartTime.Format(timeFormat)+"\n"
		endTime := "OnCalendar=*-*-* "+periodTimes.EndTime.Format(timeFormat)+"\n"

		if _, exists := uniqueTimes[startTime]; !exists {
			file.WriteString(startTime)
			uniqueTimes[startTime] = true
		}
		if _, exists := uniqueTimes[endTime]; !exists {
			file.WriteString(endTime)
			uniqueTimes[endTime] = true
		}
	}

	file.WriteString("\n[Install]\nWantedBy=timers.target")

	defer file.Close()
}

func writeService(){
	file, err := os.Create("/etc/systemd/system/sunset.service")
	if err != nil {
		fmt.Println(err)
	}
	serviceContent := `[Unit]
Description=Apply screen filter

[Service]
Type=oneshot
ExecStart=/usr/bin/sunset apply
`
	file.WriteString(serviceContent)
	defer file.Close()
}

func WriteUnitFiles(config *Config){
	writeTimer(config)
	writeService()
}

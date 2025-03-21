package main

import (
	"fmt"
	"sunset/src"
)

func main() {
	config, err := src.ParseConfig("config.toml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
                                                                 
	fmt.Println("Default Temperature:", config.Default.Temperature)
	for _, period := range config.Periods {
		fmt.Printf("Period: %dK, Start: %s, End: %s\n",
			period.Temperature,
			period.StartTime.Format("15:04:05"),
			period.EndTime.Format("15:04:05"),
		)
	}
}
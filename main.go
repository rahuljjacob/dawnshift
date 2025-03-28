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

	src.WriteUnitFiles(config)
}
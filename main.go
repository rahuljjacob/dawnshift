package main

import (
	"fmt"
	"os"
	"os/exec"
	"sunset/helpers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sunset <command>")
		fmt.Println("Commands: install, apply")
		fmt.Println("Run 'sunset --help' for help")
		return
	}

	command := os.Args[1]

	switch command {
	case "install":
		Install()
	case "apply":
		ApplyFilter()
	case "--help":
		fmt.Println("Usage: sunset <command>")
		fmt.Println("Commands:")
		fmt.Println("  install  - Install systemd service")
		fmt.Println("  apply    - Apply screen filter")
		fmt.Println("  --help   - Show this help message")
	default:
		fmt.Printf("No such command \"%s\"\n", command)
		fmt.Println("Usage: sunset <command>")
		fmt.Println("Commands: install, apply")
		fmt.Println("Run 'sunset --help' for help")
	}
}

func Install() {
	config, err := helpers.ParseConfig("./config.toml")
	if err != nil {
		fmt.Println(err)
	}
	helpers.WriteUnitFiles(config)
	cmd := exec.Command("hyprsunset")
	cmd.Run()
	helpers.ApplyFilter(config)
}

func ApplyFilter() {
	config, err := helpers.ParseConfig("./config.toml")
	if err != nil {
		fmt.Println(err)
	}
	helpers.ApplyFilter(config)
}
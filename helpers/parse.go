package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Period struct {
	Temperature int       `toml:"temperature"`
	StartTime   time.Time `toml:"start_time"`
	EndTime     time.Time `toml:"end_time"`
}

type Config struct {
	Default struct {
		Temperature int `toml:"temperature"`
	} `toml:"default"`
	Periods []Period `toml:"period"`
}

// func printConfig(config *Config){
// 	fmt.Println(config.Default.Temperature)
// 	for _, i := range config.Periods {
// 		fmt.Println(i.StartTime)
// 		fmt.Println(i.EndTime)
// 		fmt.Println(i.Temperature)
// 	}
// }

func ParseConfig() (*Config, error) {
	var config Config

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return &config, err
	}
	filePath := homeDir + "/.config/dawnshift/dawnshift.toml"

	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, fmt.Errorf("error decoding TOML: %w", err)
	}
	if config.Default.Temperature == 0 {
		config.Default.Temperature = 6250
	}
	return &config, nil
}

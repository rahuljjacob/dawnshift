package src

import (
	"fmt"
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

func ParseConfig(filePath string) (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, fmt.Errorf("error decoding TOML: %w", err)
	}

	return &config, nil
}

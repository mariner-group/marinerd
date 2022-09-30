package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/mariner-group/marinerd/internal/pkg/apiserver"
)

// Config marinerd main config struct
type Config struct {
	APIServer apiserver.APIServer `yaml:"apiServer"`
}

// Load load marinerd config
func Load(filePath string) (*Config, error) {
	if filePath == "" {
		err := errors.New("invalid config file path")
		fmt.Printf("[ERROR] %v\n", err)
		return nil, err
	}

	fmt.Printf("[INFO] load config from %v\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return nil, err
	}

	return config, nil
}

package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// ProdocsConfig is the config schema for the main prodocs application
type ProdocsConfig struct {
	Port         string `yaml:"port"`
	Repositories []struct {
		Url string `yaml:"url"`
		PAT string `yaml:"PAT"`
	} `yaml:"repositories"`
	GitFetchDuration string `yaml:"gitFetchDuration"`
}

// NewProdocsConfig initialises a new config object from a YAML file
func NewProdocsConfig(configPath string) (*ProdocsConfig, error) {
	err := validateConfigPath(configPath)
	if err != nil {
		return nil, err
	}

	config := &ProdocsConfig{}

	var file *os.File
	file, err = os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)
	if err = d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// validateConfigPath checks if the path is a readable file or not
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

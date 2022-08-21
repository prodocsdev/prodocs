/*
Copyright 2022 The ProDocs Authors. All rights reserved.
Use of this source code is governed by a BSD-style license
that can be found in the LICENSE file.
*/

package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// ProdocsConfig is the config schema for the main prodocs application
type ProdocsConfig struct {
	Port         string `yaml:"port"`
	PAT          string `yaml:"pat"`
	Repositories []struct {
		Url string `yaml:"url"`
		PAT string `yaml:"pat"`
	} `yaml:"repositories"`
	GitFetchDuration string `yaml:"gitFetchDuration"`
	StoragePath      string `yaml:"storagePath"`
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

	err = ensureStorageDir(config.StoragePath)
	if err != nil {
		log.Println("Error validating storage path: ", err, " - using default")
		config.StoragePath = os.TempDir() + "/prodocs"
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

func ensureStorageDir(dirName string) error {
	err := os.Mkdir(dirName, 0777)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

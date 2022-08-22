/*
Copyright 2022 The ProDocs Authors.
All rights reserved. Use of this source code is governed by a
BSD-style license that can be found in the LICENSE file.
*/

package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

type Repository struct {
	Url string `yaml:"url" json:"url"`
	PAT string `yaml:"PAT" json:"PAT"`
}

// ProdocsConfig is the config schema for the main prodocs application
type ProdocsConfig struct {
	// Port is the HTTP port at which the prodocs service will be available
	Port string `yaml:"port" json:"port"`

	// PAT is the GitHub PAT used to authenticate when cloning repositories
	PAT string `yaml:"PAT" json:"PAT"`

	// Repositories is the list of git repositories that need to be cloned
	Repositories []Repository `yaml:"repositories" json:"repositories"`

	// StoragePath is path at which the cloned repositories are stored
	StoragePath string `yaml:"storagePath" json:"storagePath"`

	// OutputPath is the path at which the output html files will be stored
	OutputPath string `yaml:"outputPath" json:"outputPath"`
}

// NewProdocsConfig initialises a new config object from a YAML file
func NewProdocsConfig(path string) (*ProdocsConfig, error) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		log.Println("Falling back to defaults")
	}

	config := &ProdocsConfig{
		Port:        "8080",
		StoragePath: "$HOME/prodocs/repos",
		OutputPath:  "$HOME/.prodocs/html",
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	var configJsonStr []byte
	configJsonStr, err = json.Marshal(&config)
	log.Printf("Initialised with config: %v", string(configJsonStr))

	return config, nil
}

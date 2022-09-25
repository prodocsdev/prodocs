/*
 * Copyright 2022 The ProDocs Authors.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package config

import (
	"github.com/spf13/viper"
	"log"
)

type ProdocsProjectConfig struct {
	Name   string `yaml:"name"`
	Url    string `yaml:"url"`
	Author struct {
		Name  string
		Email string
	} `yaml:"author"`
	IncludePrefixes []string `yaml:"includePrefixes"`
	ExcludePrefixes []string `yaml:"excludePrefixes"`
}

func NewProdocsProjectConfig(path string) (*ProdocsProjectConfig, error) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	config := ProdocsProjectConfig{
		IncludePrefixes: []string{"**/*.md"},
		ExcludePrefixes: []string{},
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

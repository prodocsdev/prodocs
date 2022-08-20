/*
Copyright 2022 The ProDocs Authors. All rights reserved.
Use of this source code is governed by a BSD-style license
that can be found in the LICENSE file.
*/

package cmd

import (
	"fmt"
	"github.com/prodocsdev/prodocs/config"
	"github.com/prodocsdev/prodocs/fetcher"
	"github.com/spf13/cobra"
)

var (
	path string

	fetchCommand = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch All repositories in config",
		Long:  `Clones all repositories in config, and stores them in the local repository cache.`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println(args)

			configFilePath := cmd.Flag("path").Value.String()
			config, err := config.NewProdocsConfig(configFilePath)
			if err != nil {
				panic(err)
			}
			fetcher.FetchPackages(config)
		},
	}
)

func init() {
	fetchCommand.Flags().StringVarP(&path, "path", "p", "", "config file path")
}

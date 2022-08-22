/*
Copyright 2022 The ProDocs Authors. All rights reserved.
Use of this source code is governed by a BSD-style license
that can be found in the LICENSE file.
*/

package cmd

import (
	"github.com/prodocsdev/prodocs/api"
	"github.com/prodocsdev/prodocs/config"
	"github.com/spf13/cobra"
)

var (
	path          string
	prodocsConfig *config.ProdocsConfig

	startCommand = &cobra.Command{
		Use:   "start",
		Short: "Start the prodocs service",
		Run: func(cmd *cobra.Command, args []string) {
			configFilePath := cmd.Flag("path").Value.String()

			var err error
			prodocsConfig, err = config.NewProdocsConfig(configFilePath)
			if err != nil {
				panic(err)
			}

			var server *api.ProdocsHTTPServer
			server, err = api.NewProdocsHTTPServer(prodocsConfig)
			server.Start()
		},
	}
)

func init() {
	startCommand.PersistentFlags().StringVarP(&path, "path", "p", "", "config file path")
}

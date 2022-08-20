/*
Copyright 2022 The ProDocs Authors. All rights reserved.
Use of this source code is governed by a BSD-style license
that can be found in the LICENSE file.
*/

package fetcher

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/prodocsdev/prodocs/config"
	"log"
	"os"
)

func FetchPackages(config *config.ProdocsConfig) {
	var successCountRepo = 0
	var pat string
	for _, repository := range config.Repositories {
		if repository.PAT != "" {
			pat = repository.PAT
		} else {
			pat = config.PAT
		}
		_, err := git.PlainClone("./output", false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "foo",
				Password: pat,
			},
			URL:      repository.Url,
			Progress: os.Stdout,
		})

		if err != nil {
			log.Println("Error cloning repository: ", err)
		} else {
			successCountRepo += 1
		}
	}
	log.Printf("Successfully cloned %d repositories", successCountRepo)
}

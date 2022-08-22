/*
Copyright 2022 The ProDocs Authors. All rights reserved.
Use of this source code is governed by a BSD-style license
that can be found in the LICENSE file.
*/

package renderer

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	mdparser "github.com/gomarkdown/markdown/parser"
	parser "github.com/prodocsdev/prodocs/parser"
	"os"
)

type Renderer struct {
	outputRoot string
}

func NewRenderer(outputRoot string) (*Renderer, error) {
	return &Renderer{
		outputRoot: outputRoot,
	}, nil
}

func (r *Renderer) Render(parsedMd *parser.ProdocsParsedMarkdown) error {
	extensions := mdparser.CommonExtensions | mdparser.AutoHeadingIDs
	extendedMarkdownParser := mdparser.NewWithExtensions(extensions)

	html := markdown.ToHTML(parsedMd.Data, extendedMarkdownParser, nil)
	err := os.WriteFile(fmt.Sprintf("%v/%v", r.outputRoot, parsedMd.OutputPath), html, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

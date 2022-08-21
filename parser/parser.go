package parser

import "os"

// MarkdownParser is the component that will read markdown files and find out
// template and other plugin references
type MarkdownParser struct{}

type ProdocsParsedMarkdown struct {
	Data       []byte
	OutputPath string
}

func NewMarkdownParser() (*MarkdownParser, error) {
	return &MarkdownParser{}, nil
}

func (mp *MarkdownParser) Parse(path string) (*ProdocsParsedMarkdown, error) {
	mdData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &ProdocsParsedMarkdown{
		Data: mdData,
	}, nil
}

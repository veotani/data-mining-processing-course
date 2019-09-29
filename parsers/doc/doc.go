package doc

import (
	"fmt"
	"os"

	"code.sajari.com/docconv"
)

// Parser parses doc file
type Parser struct {
	FilePath string // file to parse
}

// ExtractAllText extracts all text from file by it's file path
func (d *Parser) ExtractAllText() (string, error) {
	file, err := os.Open(d.FilePath)
	if err != nil {
		return "", err
	}
	content, _, err := docconv.ConvertDoc(file)

	if err != nil {
		return "", err
	}
	fmt.Println(content)
	return content, nil
}

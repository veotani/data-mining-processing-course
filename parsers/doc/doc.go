package doc

import (
	"os"

	"code.sajari.com/docconv"
)

// Parser parses doc file
type Parser struct {
	FilePath string // file to parse
}

// GetText extracts all text from given doc file
func (d *Parser) GetText() (string, error) {
	file, err := os.Open(d.FilePath)
	if err != nil {
		return "", err
	}
	content, _, err := docconv.ConvertDoc(file)

	if err != nil {
		return "", err
	}
	return content, nil
}

// GetMeta extracts metadata of the doc file
func (d *Parser) GetMeta() (map[string]string, error) {
	file, err := os.Open(d.FilePath)
	if err != nil {
		return nil, err
	}
	_, meta, err := docconv.ConvertDoc(file)

	if err != nil {
		return nil, err
	}
	return meta, nil
}

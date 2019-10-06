package pdf

import (
	"os"

	"code.sajari.com/docconv"
)

// Parser parses pdf file
type Parser struct {
	FileName string // file to parse
}

// GetText extracts all text from given pdf file
func (d *Parser) GetText() (string, error) {
	file, err := os.Open(d.FileName)
	if err != nil {
		return "", err
	}
	content, _, err := docconv.ConvertPDF(file)

	if err != nil {
		return "", err
	}
	return content, nil
}

// GetMeta extracts metadata of the pdf file
func (d *Parser) GetMeta() (map[string]string, error) {
	file, err := os.Open(d.FileName)
	if err != nil {
		return nil, err
	}
	_, meta, err := docconv.ConvertPDF(file)

	if err != nil {
		return nil, err
	}
	return meta, nil
}
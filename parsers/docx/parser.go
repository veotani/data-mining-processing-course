package docx

import (
	"os"

	"code.sajari.com/docconv"
)

// Parser of docx files
type Parser struct {
	FileName string
}

// GetText returns text of file without any markdown
func (p Parser) GetText() (string, error) {
	// Read file
	file, err := os.Open(p.FileName)
	if err != nil {
		return "", err
	}

	// Extract data
	text, _, err := docconv.ConvertDocx(file)
	if err != nil {
		return "", err
	}
	return text, nil
}

// GetMeta is used to extract docx document metadata
func (p Parser) GetMeta() (map[string]string, error) {
	// Read file
	file, err := os.Open(p.FileName)
	if err != nil {
		return nil, err
	}

	_, meta, err := docconv.ConvertDocx(file)
	if err != nil {
		return nil, err
	}
	return meta, nil
}

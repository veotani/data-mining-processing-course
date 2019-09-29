package doc

import (
	"errors"
	"os"
	"strings"

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
	return content, nil
}

// ExtractStudents parses doc file with table and extracts student names
func (d *Parser) ExtractStudents() ([]string, error) {
	text, err := d.ExtractAllText()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(text, "\n")
	names := make([]string, 0)

	firstNameLineNumber := -1

	for lineNumber, line := range lines {
		if strings.Contains(line, "Фамилия, Имя, Отчество") {
			firstNameLineNumber = lineNumber
			break
		}
	}

	if firstNameLineNumber == -1 {
		return nil, errors.New("cant find name line")
	}

	currentNameLineNumber := firstNameLineNumber

	for currentNameLineNumber < len(lines) && len(lines[currentNameLineNumber]) >= 6 {
		names = append(names, lines[currentNameLineNumber])
		currentNameLineNumber += 6
	}

	return names, nil
}

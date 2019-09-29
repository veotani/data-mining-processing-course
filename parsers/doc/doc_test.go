package doc

import (
	"strings"
	"testing"
)

func TestDocumentParsing(t *testing.T) {
	p := Parser{"../data/test.doc"}
	text, err := p.ExtractAllText()
	if err != nil {
		t.Error(err)
	}

	lines := strings.Split(text, "\n")

	if lines[2][:9] != "TestStart" {
		t.Errorf("first line is invalid: %v", lines[2])
	}

	if lines[len(lines)-5][:7] != "TestEnd" {
		t.Errorf("last line is invalid: %v", lines[len(lines)-5])
	}
}

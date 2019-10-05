package docx

import (
	"strings"
	"testing"
)

func TestParseEmptyFile(t *testing.T) {
	parser := Parser{"../data/empty.docx"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}

	mustBeEmpty := strings.ReplaceAll(text, " ", "")
	if mustBeEmpty != "" {
		t.Errorf("after replacing all the whitespaces content had to be empty, but it was:\n %v", mustBeEmpty)
	}
}

package pdf

import (
	"strings"
	"testing"
)

func TestEmptyFile(t *testing.T) {
	parser := Parser{FileName: "../data/empty.pdf"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}

	afterWhiteSpaceRemoval := strings.ReplaceAll(text, " ", "")
	afterNewlineRemoval := strings.ReplaceAll(afterWhiteSpaceRemoval, "\n", "")
	afterSpecialCharacterRemoval := strings.ReplaceAll(afterNewlineRemoval, "\r", "")

	if afterSpecialCharacterRemoval != "" {
		errMsg := "after replacing all whitespaces must get empty string"
		t.Errorf("%v: got string \"%q\"", errMsg, afterSpecialCharacterRemoval)
	}
}

func TestInvalidFileName(t *testing.T) {
	parser := Parser{FileName: "testBadFileName"}
	_, err := parser.GetText()
	if err == nil {
		t.Error("must throw error")
	}
}

func TestInvalidFileFormat(t *testing.T) {
	parser := Parser{FileName: "../data/invalid_pdf.pdf"}
	_, err := parser.GetText()
	if err == nil {
		t.Error("must throw error")
	}
}

func TestSimpleFile(t *testing.T) {
	mustBePhrase := "test-test-test"

	parser := Parser{FileName: "../data/simple.pdf"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}

	afterWhiteSpaceRemoval := strings.ReplaceAll(text, " ", "")
	afterNewlineRemoval := strings.ReplaceAll(afterWhiteSpaceRemoval, "\n", "")
	afterSpecialCharacterRemoval := strings.ReplaceAll(afterNewlineRemoval, "\r", "")

	if afterSpecialCharacterRemoval != mustBePhrase {
		errMsg := "after replacing all whitespaces must get \"" + mustBePhrase + "\""
		t.Errorf("%v: got string \"%v\"", errMsg, afterSpecialCharacterRemoval)
	}
}

func TestEmptyFileMeta(t *testing.T) {
	parser := Parser{FileName: "../data/empty.pdf"}
	meta, err := parser.GetMeta()
	if err != nil {
		t.Error(err)
	}
	size, ok := meta["File size"]
	if !ok {
		t.Error("file size not found")
	}

	if size != "25681 bytes" {
		t.Errorf("invalid size: %v", size)
	}
}

func TestInvalidFileFormatMeta(t *testing.T) {
	parser := Parser{FileName: "../data/invalid_pdf.pdf"}
	_, err := parser.GetMeta()
	if err == nil {
		t.Error("must throw error")
	}
}

func TestInvalidFileNameMeta(t *testing.T) {
	parser := Parser{FileName: "testBadFileName"}
	_, err := parser.GetMeta()
	if err == nil {
		t.Error("must throw error")
	}
}

package pdf

import (
	"strings"
	"testing"
)

func TestEmptyFile(t *testing.T) {
	parser := Parser{FileName: "../data/empty.doc"}
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
	parser := Parser{FileName: "../data/invalid_doc.doc"}
	_, err := parser.GetText()
	if err == nil {
		t.Error("must throw error")
	}
}

func TestSimpleFile(t *testing.T) {
	mustBePhrase := "test-test-test"

	parser := Parser{FileName: "../data/simple.doc"}
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
	parser := Parser{FileName: "../data/empty.doc"}
	meta, err := parser.GetMeta()
	if err != nil {
		t.Error(err)
	}
	if len(meta) != 0 {
		t.Error("there should be no meta")
	}
}

func TestInvalidFileFormatMeta(t *testing.T) {
	parser := Parser{FileName: "../data/invalid_doc.doc"}
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

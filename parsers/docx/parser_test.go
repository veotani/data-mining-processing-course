package docx

import (
	"strings"
	"testing"
)

func TestEmptyFile(t *testing.T) {
	parser := Parser{FileName: "../data/empty.docx"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}

	afterWhiteSpaceRemoval := strings.ReplaceAll(text, " ", "")
	afterNewlineRemoval := strings.ReplaceAll(afterWhiteSpaceRemoval, "\n", "")

	if afterNewlineRemoval != "" {
		errMsg := "after replacing all whitespaces must get empty string"
		t.Errorf("%v: got string \"%v\"", errMsg, afterNewlineRemoval)
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
	parser := Parser{FileName: "../data/invalid_docx.docx"}
	_, err := parser.GetText()
	if err == nil {
		t.Error("must throw error")
	}
}

func TestSimpleFile(t *testing.T) {
	mustBePhrase := "test-test-test"

	parser := Parser{FileName: "../data/simple.docx"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}

	afterWhiteSpaceRemoval := strings.ReplaceAll(text, " ", "")
	afterNewlineRemoval := strings.ReplaceAll(afterWhiteSpaceRemoval, "\n", "")

	if afterNewlineRemoval != mustBePhrase {
		errMsg := "after replacing all whitespaces must get \"" + mustBePhrase + "\""
		t.Errorf("%v: got string \"%v\"", errMsg, afterNewlineRemoval)
	}
}

func TestEmptyFileMeta(t *testing.T) {
	parser := Parser{FileName: "../data/empty.docx"}
	meta, err := parser.GetMeta()
	if err != nil {
		t.Error(err)
	}
	if created, ok := meta["created"]; ok != true {
		t.Errorf("file is created at 2019-10-05T17:11:50Z, but no such meta")
	} else if created != "2019-10-05T17:11:50Z" {
		t.Errorf("file is created at 2019-10-05T17:11:50Z, but got %v", created)
	}
}

func TestInvalidFileFormatMeta(t *testing.T) {
	parser := Parser{FileName: "../data/invalid_docx.docx"}
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

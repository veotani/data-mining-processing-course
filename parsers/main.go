package main

import (
	"fmt"

	"github.com/veotani/data-mining-processing-course/parsers/docx"
)

func main() {
	parser := docx.Parser{FileName: "data/test.docx"}
	text, err := parser.GetText()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(text)

	meta, err := parser.GetMeta()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(meta)
}

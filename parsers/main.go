package main

import (
	"fmt"

	"github.com/veotani/data-mining-processing-course/parsers/doc"
)

func main() {
	p := doc.Parser{FilePath: "data/test.doc"}
	text, err := p.ExtractAllText()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(text)
}

package html

import (
	"strings"
	"testing"
)

func TestApmathParse(t *testing.T) {
	parser := Parser{"http://www.apmath.spbu.ru/ru/misc/news2018.html"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(text, "Уважаемые родители и бабушки-дедушки (члены ПРОФСОЮЗА)!") {
		t.Error("wrong page was recieved")
	}
	if !strings.Contains(text, "Материалы конференции проиндексированы в РИНЦ.") {
		t.Error("recieved text doesn't contain hyperlink element")
	}
}

func TestMedium(t *testing.T) {
	parser := Parser{"https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(text, "Over the past few weeks I’ve responded to a number of questions/assertions "+
		"about pointers as a performance optimization. It seems to confuse many people, which is understandable "+
		"as it’s a complex subject. I hope this post will help.") {
		t.Error("recieved text doesn't contain first paragraph of the article")
	}
}

func TestGolangSite(t *testing.T) {
	parser := Parser{"https://golang.org"}
	text, err := parser.GetText()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(text, "Go is an open source programming language that makes it easy to build") {
		t.Error("recieved text doesn't contain information from front page")
	}
}

func TestNotALink(t *testing.T) {
	parser := Parser{"notalink"}
	text, err := parser.GetText()
	if err == nil {
		t.Errorf("error is not raised, received following text:\n%v", text)
	}
}

func TestApmathMeta(t *testing.T) {
	parser := Parser{"http://www.apmath.spbu.ru/"}
	meta, err := parser.GetMeta()
	if err != nil {
		t.Error(err)
	}

	contentStyleType, ok := meta["Content-Style-Type"]
	switch {
	case !ok:
		t.Error("Content-Style-Type not found in the meta tags")
	case contentStyleType != "text/css":
		t.Errorf("Content-Style-Type has to be \"text/css\", but it is \"%v\"", contentStyleType)
	}

	keywords, ok := meta["keywords"]
	switch {
	case !ok:
		t.Error("keywords not found in the meta tags")
	case keywords != "ПМ-ПУ, СПбГУ":
		t.Errorf("keywords has to be \"ПМ-ПУ, СПбГУ\", but it is \"%v\"", keywords)
	}

	description, ok := meta["description"]
	switch {
	case !ok:
		t.Error("description not found in the meta tags")
	case description != "Факультет прикладной математики-процессов управления. Главная страница":
		t.Errorf("description has to be \"Факультет прикладной математики-процессов управления. Главная страница\", but it is \"%v\"", description)
	}
}

func TestMediumMeta(t *testing.T) {
	parser := Parser{"https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85"}
	meta, err := parser.GetMeta()
	if err != nil {
		t.Error(err)
	}

	ogURL, ok := meta["twitter:card"]
	switch {
	case !ok:
		t.Error("no twitter:card attribute in meta tags")
	case ogURL != "summary":
		t.Errorf("twitter:card have to be \"summary\" but it is %v", ogURL)
	}
}

func TestGolangSiteMeta(t *testing.T) {
	parser := Parser{"https://golang.org"}
	meta, err := parser.GetMeta()
	if err != nil {
		t.Error(err)
	}

	description, ok := meta["description"]
	switch {
	case !ok:
		t.Error("no description attribute in meta tags")
	case description != "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.":
		t.Errorf("description have to be \"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.\" but it is %v", description)
	}
}

func TestNotALinkMeta(t *testing.T) {
	parser := Parser{"notalink"}
	meta, err := parser.GetMeta()
	if err == nil {
		t.Errorf("error is not raised, received following meta:\n%v", meta)
	}
}

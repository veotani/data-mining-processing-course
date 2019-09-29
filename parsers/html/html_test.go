package html

import (
	"errors"
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// Test cases
func Test2018YearApmathPageParser(t *testing.T) {
	url := "http://www.apmath.spbu.ru/ru/misc/news2018.html"
	posts := Parse(url)
	got := strings.Split(posts[0].text, "\n")[0]
	if got != "Уважаемые родители и бабушки-дедушки (члены ПРОФСОЮЗА)!" {
		t.Errorf("Got %v", got)
		t.Fail()
	}
}
func TestHTMLElementsParser(t *testing.T) {
	htmlPiece := `<html><div id="content"><p class="newsdate"></p><p>test</p><p class="signature"></p></div></html>`
	r := strings.NewReader(htmlPiece)
	doc, err := htmlquery.Parse(r)
	if err != nil {
		t.Errorf("There is problem with parsing hardcoded HTML: %v", err)
	}
	list := htmlquery.Find(doc, "//*[@id=\"content\"]/p|//*[@id=\"content\"]/ul")
	posts := elementsToPosts(list)

	if len(posts) != 1 {
		t.Errorf("Got %v posts. Should have gotten 1.", len(posts))
	}

	if posts[0].text != "test\n" {
		t.Errorf("Got \"%v\" as element body. Should have gotten \"test\".", posts[0].text)
	}
}

func TestHTMLElementParserWithoutPosts(t *testing.T) {
	htmlPiece := `<html><div id="content"><p class="newsdate"></p><p class="signature"></p></div></html>`
	r := strings.NewReader(htmlPiece)
	doc, err := htmlquery.Parse(r)
	if err != nil {
		t.Errorf("There is problem with parsing hardcoded HTML: %v", err)
	}
	list := htmlquery.Find(doc, "//*[@id=\"content\"]/p|//*[@id=\"content\"]/ul")
	posts := elementsToPosts(list)

	if len(posts) != 1 {
		t.Errorf("Got %v posts. Should have gotten 0.", len(posts))
	}

	if posts[0].text != "" {
		t.Errorf("Got \"%v\" as text content. Should have gotten \"\".", posts[0].text)
	}
}

// Unit tests
func TestGetAttr(t *testing.T) {
	htmlPiece := `<div id="content"></div>`
	r := strings.NewReader(htmlPiece)
	doc, err := htmlquery.Parse(r)
	if err != nil {
		t.Errorf("There is problem with parsing hardcoded HTML: %v", err)
	}

	div := doc.FirstChild.LastChild.FirstChild

	attr, err := getAttr(div, "id")

	if err != nil {
		t.Errorf("Expected attribute name but attribute wasn't found.")
	}

	if attr != "content" {
		t.Errorf("Expected: \"content\" but got \"%v\"", attr)
	}
}

func TestElementsToPosts(t *testing.T) {
	htmlElements := make([]*html.Node, 3)

	elem, err := buildElementByString("<p></p>")
	if err != nil {
		t.Fail()
	}
	htmlElements[0] = elem
	elem, err = buildElementByString("<p class=\"signature\"></p>")
	if err != nil {
		t.Fail()
	}
	htmlElements[1] = elem
	elem, err = buildElementByString("<p class=\"newsdate\"></p>")
	if err != nil {
		t.Fail()
	}
	htmlElements[2] = elem

	posts := elementsToPosts(htmlElements)

	if len(posts) != 1 {
		t.Errorf("Posts len must be 1. Got %v", len(posts))
	}

	if posts[0].text != "" {
		t.Errorf("Text body must be empty, but got %v", posts[0].text)
	}
}

func buildElementByString(element string) (*html.Node, error) {
	elem, err := htmlquery.Parse(strings.NewReader(element))
	if err != nil {
		return nil, errors.New("Can't build element")
	}
	return elem.FirstChild.LastChild.FirstChild, nil
}

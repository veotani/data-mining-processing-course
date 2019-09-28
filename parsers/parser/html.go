package parser

import (
	"errors"
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Post struct {
	text string
}

// Parse returns all apmath faculty web-site posts as text
func Parse() []*Post {
	fmt.Println("Extracting news:")
	elements := extractNewsElements("http://apmath.spbu.ru/")
	posts := elementsToPosts(elements)
	return posts
}

func extractNewsElements(url string) []*html.Node {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println(err)
	}
	list := htmlquery.Find(doc, "//*[@id=\"content\"]/p|//*[@id=\"content\"]/ul")
	return list
}

func elementsToPosts(elements []*html.Node) []*Post {
	countElements := len(elements)
	currentElement := 0
	result := make([]*Post, 0)

	for currentElement < countElements {
		class, err := getAttr(elements[currentElement], "class")
		if err != nil || class != "newsdate" {
			currentElement++
			continue
		}

		currentElement++

		text := ""
		for currentElement < countElements {
			class, _ := getAttr(elements[currentElement], "class")
			if class == "signature" {
				currentElement++
				break
			}
			text += htmlquery.InnerText(elements[currentElement]) + "\n"
			currentElement++
		}
		result = append(result, &Post{text})
	}

	return result
}

func getAttr(el *html.Node, attrName string) (string, error) {
	if el.Attr == nil {
		return "", errors.New("There is no such attribute")
	}

	for _, a := range el.Attr {
		if a.Key == attrName {
			return a.Val, nil
		}
	}

	return "", errors.New("There is no such attribute")
}

package html

import (
	nethtml "golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
)

// Parser object in html package used to extract text from html page by it's URL
type Parser struct {
	URL string
}

// GetText method extracts text of page by URL in the Parser object
func (p Parser) GetText() (string, error) {
	doc, err := htmlquery.LoadURL(p.URL)
	if err != nil {
		return "", err
	}
	return htmlquery.InnerText(doc), nil
}

// GetMeta is used to extract html page metadata
func (p Parser) GetMeta() (map[string]string, error) {
	doc, err := htmlquery.LoadURL(p.URL)
	if err != nil {
		return nil, err
	}

	meta := make(map[string]string)

	metaTags := htmlquery.Find(doc, "/html/head/meta")
	for _, metaTag := range metaTags {
		key, value := parseMetaTagAttributes(metaTag.Attr)
		if key != "" || value != "" {
			meta[key] = value
		}
	}

	return meta, nil
}

func parseMetaTagAttributes(attributes []nethtml.Attribute) (key string, value string) {
	key, value = "", ""
	for _, attr := range attributes {
		if attr.Key == "charset" {
			key, value = attr.Key, attr.Val
		}
		if attr.Key == "http-equiv" || attr.Key == "name" {
			key = attr.Val
		}
		if attr.Key == "content" {
			value = attr.Val
		}
	}
	return
}

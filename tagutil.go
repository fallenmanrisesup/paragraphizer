package paragraphizer

import (
	"fmt"

	"golang.org/x/net/html"
)

var (
	inlineTags = map[string]bool{
		"a":      true,
		"abbr":   true,
		"b":      true,
		"br":     true,
		"cite":   true,
		"code":   true,
		"em":     true,
		"i":      true,
		"img":    true,
		"kbd":    true,
		"mark":   true,
		"q":      true,
		"ruby":   true,
		"s":      true,
		"script": true,
		"small":  true,
		"strong": true,
		"sub":    true,
		"sup":    true,
		"span":   true,
		"time":   true,
		"u":      true,
		"var":    true,
		"wbr":    true,
	}

	blockTags = map[string]bool{
		"address":  true,
		"article":  true,
		"aside":    true,
		"div":      true,
		"dl":       true,
		"dt":       true,
		"dd":       true,
		"footer":   true,
		"form":     true,
		"h1":       true,
		"h2":       true,
		"h3":       true,
		"h4":       true,
		"h5":       true,
		"h6":       true,
		"header":   true,
		"hgroup":   true,
		"hr":       true,
		"main":     true,
		"markdown": true,
		"nav":      true,
		"ol":       true,
		"p":        true,
		"pre":      true,
		"section":  true,
		"table":    true,
		"tbody":    true,
		"td":       true,
		"th":       true,
		"thead":    true,
		"tr":       true,
		"ul":       true,
	}
)

func isAllowedInline(tag string) bool {

	return inlineTags[tag]
}

func isBlockTag(tag string) bool {

	return blockTags[tag]
}

func formatStartTag(token html.Token) string {
	attrs := ""
	for _, attr := range token.Attr {
		attrs += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}
	return fmt.Sprintf("<%s%s>", token.Data, attrs)
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	var buf bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(&buf)
	if err != nil {
		log.Fatal(err)
	}

	processHTMLNode(doc, 0)
}

func processHTMLNode(n *html.Node, nesting int) {
	// filter
	if n.Type == html.TextNode {
		if strings.TrimSpace(n.Data) == "" {
			processNestingHTMLNode(n, nesting)
			return
		}
	}
	if n.Type == html.DocumentNode {
		processNestingHTMLNode(n, nesting)
		return
	}
	if n.Type == html.ElementNode {
		switch n.DataAtom {
		case atom.Html, atom.Head, atom.Body:
			processNestingHTMLNode(n, nesting)
			return
		}
	}

	// output
	outputHTMLNode(n, nesting)

	// recursion
	processNestingHTMLNode(n, nesting)

	// output end
	outputHTMLNodeEnd(n, nesting)
}

func processNestingHTMLNode(n *html.Node, nesting int) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nesting++
		processHTMLNode(c, nesting)
		nesting--
	}
}

func indent(nesting int) string {
	var b strings.Builder
	for i := 0; i < nesting*2; i++ {
		b.WriteByte(' ')
	}
	indent := b.String()
	if len(indent) >= 6 {
		indent = indent[6:]
	}
	return indent
}

func outputHTMLNode(n *html.Node, nesting int) {
	indent := indent(nesting)

	if n.Type == html.CommentNode {
		fmt.Printf("%s// %s\n", indent, strings.TrimSpace(n.Data))
		return
	}

	if n.Type == html.ElementNode {
		outputHTMLElement(n, indent)
		return
	}

	if n.Type == html.TextNode {
		fmt.Printf("%svecty.Text(\"%s\"),\n", indent, strings.TrimSpace(n.Data))
		return
	}

	fmt.Printf("%s??? // type(%v) atom(%v) data(%v)\n", indent, n.Type, n.DataAtom, n.Data)
}

func outputHTMLNodeEnd(n *html.Node, nesting int) {
	indent := indent(nesting)

	if n.Type == html.ElementNode {
		fmt.Printf("%s),\n", indent)
	}
}

func outputHTMLElement(n *html.Node, indent string) {
	elemName := elemName(n)
	fmt.Printf("%selem.%v(\n", indent, elemName)
	if len(n.Attr) > 0 {
		fmt.Printf("  %svecty.Markup(\n", indent)
		for _, attr := range n.Attr {
			attrCode := attrCode(attr.Key, attr.Val)
			fmt.Printf("    %s%s,\n", indent, attrCode)
		}
		fmt.Printf("  %s),\n", indent)
	}
}

func elemName(n *html.Node) string {
	switch tag := n.DataAtom; tag {
	case atom.Li:
		return "ListItem"
	case atom.A:
		return "Anchor"
	case atom.I:
		return "Italic"
	case atom.Img:
		return "Image"
	case atom.H1:
		return "Heading1"
	case atom.H2:
		return "Heading2"
	case atom.H3:
		return "Heading3"
	case atom.H4:
		return "Heading4"
	case atom.H5:
		return "Heading5"
	case atom.H6:
		return "Heading6"
	default:
		return strings.Title(tag.String())
	}
}

func attrCode(key, value string) string {
	var property = map[string]bool{
		"autofocus":   true,
		"disabled":    true,
		"checked":     true,
		"htmlFor":     true,
		"href":        true,
		"id":          true,
		"placeholder": true,
		"src":         true,
		"type":        true,
		"value":       true,
		"name":        true,
		"alt":         true,
	}
	var style = map[string]bool{
		"color":      true,
		"width":      true,
		"min-width":  true,
		"max-width":  true,
		"height":     true,
		"min-height": true,
		"max-height": true,
		"margin":     true,
		"overflow":   true,
		"overflow-x": true,
		"overflow-y": true,
	}
	var b strings.Builder

	if key == "class" {
		fmt.Fprintf(&b, "vecty.Class(")
		for i, field := range strings.Fields(value) {
			if i > 0 {
				fmt.Fprintf(&b, ", ")
			}
			fmt.Fprintf(&b, "%q", field)
		}
		fmt.Fprintf(&b, ")")
		return b.String()
	}

	if property[key] {
		fmt.Fprintf(&b, "vecty.Property(%q, %q)", key, value)
		return b.String()
	}

	if style[key] {
		fmt.Fprintf(&b, "vecty.Style(%q, %q)", key, value)
		return b.String()
	}

	if strings.HasPrefix(key, "data-") {
		fmt.Fprintf(&b, "vecty.Data(%q, %q)", key[5:], value)
		return b.String()
	}

	fmt.Fprintf(&b, "vecty.Attribute(%q, %q)", key, value)
	return b.String()
}

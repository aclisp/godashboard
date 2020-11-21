package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
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
	// ignore document, <html>, <head> and <body>
	if nesting < 3 && (n.Type == html.DocumentNode || n.Type == html.ElementNode) {
		processNestingHTMLNode(n, nesting)
		return
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
	// remove extra indents
	if len(indent) >= 6 {
		indent = indent[6:]
	}
	return indent
}

func outputHTMLNode(n *html.Node, nesting int) {
	indent := indent(nesting)

	if n.Type == html.CommentNode {
		if text := strings.TrimSpace(n.Data); text != "" {
			fmt.Printf("%s// %s\n", indent, text)
		}
		return
	}

	if n.Type == html.ElementNode {
		outputHTMLElement(n, indent)
		return
	}

	if n.Type == html.TextNode {
		if text := strings.TrimSpace(n.Data); text != "" {
			fmt.Printf("%svecty.Text(\"%s\"),\n", indent, text)
		}
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
		fmt.Printf("  %svecty.Markup(", indent)
		moveClassToFront(n.Attr)
		for i, attr := range n.Attr {
			attrCode := attrCode(attr.Key, attr.Val)
			if i == 0 {
				fmt.Printf("%s", attrCode)
			} else {
				fmt.Printf(",\n    %s%s", indent, attrCode)
			}
		}
		if len(n.Attr) > 1 {
			fmt.Printf(",\n  %s", indent)
		}
		fmt.Printf("),\n")
	}
}

func moveClassToFront(attrs []html.Attribute) {
	// find the first `class` attribute
	var class html.Attribute
	for _, attr := range attrs {
		if attr.Key == "class" {
			class = attr
			break
		}
	}
	if class.Key == "" { // not found
		return
	}
	moveToFront(class, attrs)
}

func moveToFront(class html.Attribute, attrs []html.Attribute) {
	if len(attrs) == 0 || attrs[0] == class {
		return
	}
	var prev html.Attribute
	for i, elem := range attrs {
		switch {
		case i == 0:
			attrs[0] = class
			prev = elem
		case elem == class:
			attrs[i] = prev
			return
		default:
			attrs[i] = prev
			prev = elem
		}
	}
}

func elemName(n *html.Node) string {
	var elemNameMap = map[string]string{
		"a":          "Anchor",
		"abbr":       "Abbreviation",
		"b":          "Bold",
		"bdi":        "BidirectionalIsolation",
		"bdo":        "BidirectionalOverride",
		"blockquote": "BlockQuote",
		"br":         "Break",
		"cite":       "Citation",
		"col":        "Column",
		"colgroup":   "ColumnGroup",
		"datalist":   "DataList",
		"dd":         "Description",
		"del":        "DeletedText",
		"dfn":        "Definition",
		"dl":         "DescriptionList",
		"dt":         "DefinitionTerm",
		"em":         "Emphasis",
		"fieldset":   "FieldSet",
		"figcaption": "FigureCaption",
		"h1":         "Heading1",
		"h2":         "Heading2",
		"h3":         "Heading3",
		"h4":         "Heading4",
		"h5":         "Heading5",
		"h6":         "Heading6",
		"hgroup":     "HeadingsGroup",
		"hr":         "HorizontalRule",
		"i":          "Italic",
		"iframe":     "InlineFrame",
		"img":        "Image",
		"ins":        "InsertedText",
		"kbd":        "KeyboardInput",
		"li":         "ListItem",
		"menuitem":   "MenuItem",
		"nav":        "Navigation",
		"noframes":   "NoFrames",
		"noscript":   "NoScript",
		"ol":         "OrderedList",
		"optgroup":   "OptionsGroup",
		"p":          "Paragraph",
		"param":      "Parameter",
		"pre":        "Preformatted",
		"q":          "Quote",
		"rp":         "RubyParenthesis",
		"rt":         "RubyText",
		"rtc":        "RubyTextContainer",
		"s":          "Strikethrough",
		"samp":       "Sample",
		"sub":        "Subscript",
		"sup":        "Superscript",
		"tbody":      "TableBody",
		"textarea":   "TextArea",
		"td":         "TableData",
		"tfoot":      "TableFoot",
		"th":         "TableHeader",
		"thead":      "TableHead",
		"tr":         "TableRow",
		"u":          "Underline",
		"ul":         "UnorderedList",
		"var":        "Variable",
		"wbr":        "WordBreakOpportunity",
	}
	tagName := n.DataAtom.String()
	funName := elemNameMap[tagName]
	if funName == "" {
		funName = strings.Title(tagName)
	}
	return funName
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

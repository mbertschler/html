package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mbertschler/html/cmd/internal"
	"golang.org/x/net/html"
)

var testInput = `<!DOCTYPE html>
<html>
	<head>
		<title>Page Title</title>
	</head>
	<body>
		<h1>This is a Heading</h1>
		<p>This is a paragraph.</p>

		<p data-test="yes">Links:</p>
		<ul>
			<li>
				<a href="foo">Foo</a>
			<li>
				<a href="/bar/baz">BarBaz</a>
		</ul>
		<custom-element custom-attribute="a"></custom-element>
	</body>
</html>`

var (
	elementsMap   = map[string]bool{}
	attributesMap = map[string]bool{}
)

func initMaps() {
	for _, element := range internal.Elements {
		elementsMap[element.Name] = true
	}
	for _, attribute := range internal.Attributes {
		attributesMap[attribute.Name] = true
	}
}

func main() {
	doc, err := html.Parse(strings.NewReader(testInput))
	if err != nil {
		log.Fatal(err)
	}
	initMaps()
	handleNode(os.Stdout, doc)
}

var indentLevel int = -1

func printIndent(w io.Writer) {
	for i := 0; i < indentLevel; i++ {
		fmt.Fprintf(w, "  ")
	}
}

func handleNode(w io.Writer, n *html.Node) {
	indentLevel++

	switch n.Type {
	case html.TextNode:
		trimmed := strings.TrimSpace(n.Data)
		if trimmed != "" {
			printIndent(w)
			fmt.Fprintf(w, "html.Text(%q),\n", trimmed)
		}
	case html.DocumentNode:
		handleElement(w, n)
	case html.ElementNode:
		handleElement(w, n)
	case html.CommentNode:
		printIndent(w)
		fmt.Fprintf(w, "html.Comment(%q),\n", n.Data)
	case html.DoctypeNode:
		printIndent(w)
		fmt.Fprintf(w, "html.Doctype(%q),\n", n.Data)

	// printing nodes that are not handled
	case html.ErrorNode:
		log.Printf("ErrorNode: %s\n", n.Data)
	default:
		log.Printf("unexpected node: %s\n", n.Data)
	}
	indentLevel--
}

func handleElement(w io.Writer, n *html.Node) {
	printIndent(w)

	switch n.Type {
	case html.DocumentNode:
		fmt.Fprintf(w, "html.Blocks{\n")
	case html.ElementNode:
		writeOpenTag(w, n.Data)
		first := true
		if len(n.Attr) > 0 {
			for _, a := range n.Attr {
				if first {
					fmt.Fprintf(w, "attr")
					first = false
				}
				writeAttr(w, a)
			}
		} else {
			fmt.Fprintf(w, "nil")
		}
		if n.FirstChild != nil {
			fmt.Fprintf(w, ",\n")
		}
	default:
		log.Printf("unexpected node type: %v\n", n.Type)
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		handleNode(w, c)
	}
	switch n.Type {
	case html.DocumentNode:
		fmt.Fprintf(w, "}\n")
	case html.ElementNode:
		if n.FirstChild != nil {
			printIndent(w)
		}
		fmt.Fprintf(w, "),\n")
	}
}

func writeOpenTag(w io.Writer, name string) {
	name = strings.ToLower(name)
	ok := elementsMap[name]
	if !ok {
		fmt.Fprintf(w, "html.Element(%q, ", name)
		return
	}
	fmt.Fprintf(w, "html.%s(", strings.Title(name))
}

func writeAttr(w io.Writer, a html.Attribute) {
	name := strings.ToLower(a.Key)
	ok := attributesMap[name]
	if !ok {
		if strings.HasPrefix(name, "data-") {
			fmt.Fprintf(w, ".DataAttr(%q, %q)", name[5:], a.Val)
			return
		}
		fmt.Fprintf(w, ".Attr(%q, %q)", name, a.Val)
		return
	}
	fmt.Fprintf(w, ".%s(%q)", strings.Title(name), a.Val)
}

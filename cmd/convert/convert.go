package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

		<p>Links:</p>
		<ul>
			<li>
				<a href="foo">Foo</a>
			<li>
				<a href="/bar/baz">BarBaz</a>
		</ul>
	</body>
</html>`

func main() {
	doc, err := html.Parse(strings.NewReader(testInput))
	if err != nil {
		log.Fatal(err)
	}

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
		fmt.Fprintf(w, "html.Element(%q, ", n.Data)
		first := true
		if len(n.Attr) > 0 {
			for _, a := range n.Attr {
				if first {
					fmt.Fprintf(w, "attr")
					first = false
				}
				fmt.Fprintf(w, ".Attr(%q, %q)", a.Key, a.Val)
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

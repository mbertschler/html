package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mbertschler/html/cmd/internal"
	"golang.org/x/net/html"
)

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
	flag.Parse()
	args := flag.Args()
	var input io.Reader
	if len(args) > 0 {
		input = strings.NewReader(strings.Join(args, " "))
	} else {
		input = os.Stdin
		go func() {
			time.Sleep(time.Second)
			log.Println("waiting for input on stdin...")
		}()
	}
	doc, err := html.Parse(input)
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

//go:generate go run ./cmd/generate
package html

import "github.com/mbertschler/html/attr"

func Doctype(arg string) Block {
	return newElement("!DOCTYPE", attr.Attributes{{Key: arg}}, nil, Void)
}

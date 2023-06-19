package main

import "github.com/mbertschler/html"

type element struct {
	Name   string
	Option html.ElementOption
}

// Source: https://developer.mozilla.org/en-US/docs/Web/HTML/Element
// taken from the sidebar, deprecated elements removed
var elements = []element{
	{Name: "a"},
	{Name: "abbr"},
	{Name: "address"},
	{Name: "area"},
	{Name: "article"},
	{Name: "aside"},
	{Name: "audio"},
	{Name: "b"},
	{Name: "base"},
	{Name: "bdi"},
	{Name: "bdo"},
	{Name: "blockquote"},
	{Name: "body"},
	{Name: "br", Option: html.Void},
	{Name: "button"},
	{Name: "canvas"},
	{Name: "caption"},
	{Name: "cite"},
	{Name: "code"},
	{Name: "col"},
	{Name: "colgroup"},
	{Name: "data"},
	{Name: "datalist"},
	{Name: "dd"},
	{Name: "del"},
	{Name: "details"},
	{Name: "dfn"},
	{Name: "dialog"},
	{Name: "div"},
	{Name: "dl"},
	{Name: "dt"},
	{Name: "em"},
	{Name: "embed"},
	{Name: "fieldset"},
	{Name: "figcaption"},
	{Name: "figure"},
	{Name: "footer"},
	{Name: "form"},
	{Name: "h1"},
	{Name: "h2"},
	{Name: "h3"},
	{Name: "h4"},
	{Name: "h5"},
	{Name: "h6"},
	{Name: "head"},
	{Name: "header"},
	{Name: "hgroup"},
	{Name: "hr", Option: html.Void},
	{Name: "html"},
	{Name: "i"},
	{Name: "iframe"},
	{Name: "img", Option: html.Void},
	{Name: "input", Option: html.Void},
	{Name: "ins"},
	{Name: "kbd"},
	{Name: "label"},
	{Name: "legend"},
	{Name: "li"},
	{Name: "link", Option: html.Void},
	{Name: "main"},
	{Name: "map"},
	{Name: "mark"},
	{Name: "menu"},
	{Name: "meta", Option: html.Void},
	{Name: "meter"},
	{Name: "nav"},
	{Name: "noscript"},
	{Name: "object"},
	{Name: "ol"},
	{Name: "optgroup"},
	{Name: "option"},
	{Name: "output"},
	{Name: "p"},
	{Name: "picture"},
	{Name: "portal"},
	{Name: "pre", Option: html.NoWhitespace},
	{Name: "progress"},
	{Name: "q"},
	{Name: "rp"},
	{Name: "rt"},
	{Name: "ruby"},
	{Name: "s"},
	{Name: "samp"},
	{Name: "script", Option: html.JSElement},
	{Name: "section"},
	{Name: "select"},
	{Name: "slot"},
	{Name: "small"},
	{Name: "source"},
	{Name: "span"},
	{Name: "strong"},
	{Name: "style", Option: html.CSSElement},
	{Name: "sub"},
	{Name: "summary"},
	{Name: "sup"},
	{Name: "table"},
	{Name: "tbody"},
	{Name: "td"},
	{Name: "template"},
	{Name: "textarea", Option: html.NoWhitespace},
	{Name: "tfoot"},
	{Name: "th"},
	{Name: "thead"},
	{Name: "time"},
	{Name: "title"},
	{Name: "tr"},
	{Name: "track"},
	{Name: "u"},
	{Name: "ul"},
	{Name: "var"},
	{Name: "video"},
	{Name: "wbr"},
}

var elementsFileTemplate = `// Code generated by cmd/generate. DO NOT EDIT.
package html

import "github.com/mbertschler/html/attr"

{{range . -}}
{{ if .NoChildren -}}
func {{ .FuncName }}(attr attr.Attributes) Block {
	return newElement("{{ .TagName }}", attr, nil, Void)
}
{{ else -}}
func {{ .FuncName }}(attr attr.Attributes, children ...Block) Block {
	return newElement("{{ .TagName }}", attr, children, {{ .Option }})
}
{{ end }}{{ end }}`

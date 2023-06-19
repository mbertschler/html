package html_test

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"testing"

	"github.com/mbertschler/html"
	"github.com/mbertschler/html/attr"
)

func TestHTMLTemplate(t *testing.T) {
	buf := &bytes.Buffer{}
	renderTemplate(buf)
}

func BenchmarkHTMLTemplate(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		renderTemplate(buf)
		buf.Reset()
	}
}

func TestBlocks(t *testing.T) {
	buf := &bytes.Buffer{}
	renderBlocks(buf)

}

func BenchmarkBlocks(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		renderBlocks(buf)
		buf.Reset()
	}
}

func BenchmarkBlockTemplates(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		renderBlockTemplates(buf)
		buf.Reset()
	}
}

var t *template.Template

func renderTemplate(w io.Writer) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`
	var err error
	if t == nil {
		t, err = template.New("webpage").Parse(tpl)
		if err != nil {
			log.Fatal(err)
		}
	}

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(w, noItems)
	if err != nil {
		log.Fatal(err)
	}
}

type Placeholder struct {
	Pointer html.Block
}

func (p Placeholder) RenderHTML() html.Block {
	return p.Pointer
}

type Template struct {
	block html.Block
	title Placeholder
	rows  Placeholder
}

func (t *Template) Render(title string, rows html.Blocks) html.Block {
	t.title.Pointer = html.Text(title)
	t.rows.Pointer = rows
	return t.block
}

func NewTemplate() *Template {
	out := &Template{}
	out.block = html.Blocks{
		html.Doctype("html"),
		html.Html(nil,
			html.Head(nil,
				html.Meta(attr.Charset("UTF-8")),
				html.Title(nil, &out.title),
			),
			html.Body(nil, &out.rows),
		),
	}
	return out
}

func TestTemplate(t *testing.T) {
	template := NewTemplate()
	block := template.Render("My page", html.Blocks{})
	var out = &bytes.Buffer{}
	err := html.RenderMinified(out, block)
	if err != nil {
		log.Fatal(err)
	}
	expected := `<!DOCTYPE html><html><head><meta charset="UTF-8"><title>My page</title></head><body></body></html>`
	if out.String() != expected {
		t.Fatalf("expected %q, got %q", expected, out.String())
	}

	block = template.Render("My other page", html.Blocks{})
	out.Reset()
	err = html.RenderMinified(out, block)
	if err != nil {
		log.Fatal(err)
	}
	expected = `<!DOCTYPE html><html><head><meta charset="UTF-8"><title>My other page</title></head><body></body></html>`
	if out.String() != expected {
		t.Fatalf("expected %q, got %q", expected, out.String())
	}
}

func renderBlocks(w io.Writer) {
	type Data struct {
		Title string
		Items []string
	}

	blocks := func(d Data) html.Block {
		var rows html.Blocks
		if len(d.Items) == 0 {
			rows.Add(html.Div(nil, html.Strong(nil, html.Text("no rows"))))
		} else {
			for _, e := range d.Items {
				rows.Add(html.Div(nil, html.Text(e)))
			}
		}
		return html.Blocks{
			html.Doctype("html"),
			html.Html(nil,
				html.Head(nil,
					html.Meta(attr.Charset("UTF-8")),
					html.Title(nil, html.Text(d.Title)),
				),
				html.Body(nil, rows),
			),
		}
	}

	data := Data{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	err := html.RenderMinified(w, blocks(data))
	if err != nil {
		log.Fatal(err)
	}

	noItems := Data{
		Title: "My another page",
		Items: []string{},
	}

	err = html.RenderMinified(w, blocks(noItems))
	if err != nil {
		log.Fatal(err)
	}
}

func renderBlockTemplates(w io.Writer) {
	template := NewTemplate()
	type Data struct {
		Title string
		Items []string
	}

	blocks := func(d Data) html.Block {
		var rows html.Blocks
		if len(d.Items) == 0 {
			rows.Add(html.Div(nil, html.Strong(nil, html.Text("no rows"))))
		} else {
			for _, e := range d.Items {
				rows.Add(html.Div(nil, html.Text(e)))
			}
		}
		return template.Render(d.Title, rows)
	}

	data := Data{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	err := html.RenderMinified(w, blocks(data))
	if err != nil {
		log.Fatal(err)
	}

	noItems := Data{
		Title: "My another page",
		Items: []string{},
	}

	err = html.RenderMinified(w, blocks(noItems))
	if err != nil {
		log.Fatal(err)
	}
}

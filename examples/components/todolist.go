package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mbertschler/blocks/html"
)

// started work based on the template at 13:45
// done implementing the static template at 14:05

func todoLayout(todoApp html.Block) html.Block {
	return html.Blocks{
		html.Doctype("html"),
		html.Html(html.Attr("lang", "en"),
			html.Head(nil,
				html.Meta(html.Charset("utf-8")),
				html.Meta(html.Name("viewport").Content("width=device-width, initial-scale=1")),
				html.Title(nil, html.Text("Guiapi • TodoMVC")),
				html.Link(html.Rel("stylesheet").Href("https://cdn.jsdelivr.net/npm/todomvc-app-css@2.4.2/index.min.css")),
			),
			html.Body(nil,
				todoApp,
				html.Elem("footer", html.Class("info"),
					html.P(nil, html.Text("Double-click to edit a todo")),
					html.P(nil, html.Text("Template by "), html.A(html.Href("http://sindresorhus.com"), html.Text("Sindre Sorhus"))),
					html.P(nil, html.Text("Created by "), html.A(html.Href("https://github.com/mbertschler"), html.Text("Martin Bertschler"))),
					html.P(nil, html.Text("Part of "), html.A(html.Href("http://todomvc.com"), html.Text("TodoMVC"))),
				),
				html.Script(html.Src("/js/guiapi.js")),
				html.Script(html.Src("/js/app.js")),
			),
		),
	}
}

type TodoList struct {
	*App
}

func (t *TodoList) Component() *ComponentConfig {
	return &ComponentConfig{
		Name:    "TodoList",
		Actions: map[string]ActionFunc{},
	}
}

func (t *TodoList) RenderPage(ctx *gin.Context) (html.Block, error) {
	appBlock, err := t.renderAppBlock(ctx)
	if err != nil {
		return nil, err
	}
	return todoLayout(appBlock), nil
}

func (t *TodoList) renderAppBlock(ctx *gin.Context) (html.Block, error) {
	main := html.Elem("section", html.Class("main"),
		html.Input(html.Class("toggle-all").Attr("type", "checkbox")),
		html.Label(html.Attr("for", "toggle-all"), html.Text("Mark all as complete")),
		html.Ul(html.Class("todo-list"),
			html.Li(html.Class("completed"),
				html.Div(html.Class("view"),
					html.Input(html.Class("toggle").Attr("type", "checkbox").Attr("checked", "")),
					html.Label(nil, html.Text("Taste JavaScript")),
					html.Button(html.Class("destroy")),
				),
				html.Input(html.Class("edit").Attr("value", "Create a TodoMVC template")),
			),
			html.Li(nil,
				html.Div(html.Class("view"),
					html.Input(html.Class("toggle").Attr("type", "checkbox")),
					html.Label(nil, html.Text("Buy a unicorn")),
					html.Elem("button", html.Class("destroy")),
				),
				html.Input(html.Class("edit").Attr("value", "Rule the web")),
			),
		),
	)

	footer := html.Elem("footer", html.Class("footer"),
		html.Span(html.Class("todo-count"),
			html.Strong(nil, html.Text("2")),
			html.Text(" items left"),
		),
		html.Ul(html.Class("filters"),
			html.Li(nil,
				html.A(html.Class("selected").Attr("href", "#/"), html.Text("All")),
			),
			html.Li(nil,
				html.A(html.Attr("href", "#/active"), html.Text("Active")),
			),
			html.Li(nil,
				html.A(html.Attr("href", "#/completed"), html.Text("Completed")),
			),
		),
		html.Button(html.Class("clear-completed"), html.Text("Clear completed")),
	)

	block := html.Elem("section", html.Class("todoapp"),
		html.Elem("header", html.Class("header"),
			html.H1(nil, html.Text("todos")),
			html.Input(html.Class("new-todo").Attr("placeholder", "What needs to be done?").Attr("autofocus", "")),
		),
		main,
		footer,
	)
	return block, nil
}

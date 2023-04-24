package main

type attribute struct {
	Name  string
	IsURL bool
}

var attributes = []attribute{
	{Name: "accept"},
	{Name: "accept-charset"},
	{Name: "accesskey"},
	{Name: "action", IsURL: true},
	{Name: "align"},
	{Name: "allow"},
	{Name: "alt"},
	{Name: "async"},
	{Name: "autocapitalize"},
	{Name: "autocomplete"},
	{Name: "autofocus"},
	{Name: "autoplay"},
	{Name: "background", IsURL: true},
	{Name: "bgcolor"},
	{Name: "border"},
	{Name: "buffered"},
	{Name: "capture"},
	{Name: "challenge"},
	{Name: "charset"},
	{Name: "checked"},
	{Name: "cite", IsURL: true},
	{Name: "class"},
	{Name: "code"},
	{Name: "codebase", IsURL: true},
	{Name: "color"},
	{Name: "cols"},
	{Name: "colspan"},
	{Name: "content"},
	{Name: "contenteditable"},
	{Name: "contextmenu"},
	{Name: "controls"},
	{Name: "coords"},
	{Name: "crossorigin"},
	{Name: "data", IsURL: true},
	{Name: "data-*"},
	{Name: "datetime"},
	{Name: "decoding"},
	{Name: "default"},
	{Name: "defer"},
	{Name: "dir"},
	{Name: "dirname"},
	{Name: "disabled"},
	{Name: "download"},
	{Name: "draggable"},
	{Name: "enctype"},
	{Name: "for"},
	{Name: "form"},
	{Name: "formaction", IsURL: true},
	{Name: "formenctype"},
	{Name: "formmethod"},
	{Name: "formnovalidate"},
	{Name: "formtarget"},
	{Name: "headers"},
	{Name: "height"},
	{Name: "hidden"},
	{Name: "high"},
	{Name: "href", IsURL: true},
	{Name: "hreflang"},
	{Name: "http-equiv"},
	{Name: "id"},
	{Name: "integrity"},
	{Name: "inputmode"},
	{Name: "ismap"},
	{Name: "itemprop"},
	{Name: "keytype"},
	{Name: "kind"},
	{Name: "label"},
	{Name: "lang"},
	{Name: "list"},
	{Name: "loop"},
	{Name: "low"},
	{Name: "max"},
	{Name: "maxlength"},
	{Name: "minlength"},
	{Name: "media"},
	{Name: "method"},
	{Name: "min"},
	{Name: "multiple"},
	{Name: "muted"},
	{Name: "name"},
	{Name: "novalidate"},
	{Name: "open"},
	{Name: "optimum"},
	{Name: "pattern"},
	{Name: "ping"},
	{Name: "placeholder"},
	{Name: "playsinline"},
	{Name: "poster", IsURL: true},
	{Name: "preload"},
	{Name: "readonly"},
	{Name: "referrerpolicy"},
	{Name: "rel"},
	{Name: "required"},
	{Name: "reversed"},
	{Name: "role"},
	{Name: "rows"},
	{Name: "rowspan"},
	{Name: "sandbox"},
	{Name: "scope"},
	{Name: "selected"},
	{Name: "shape"},
	{Name: "size"},
	{Name: "sizes"},
	{Name: "slot"},
	{Name: "span"},
	{Name: "spellcheck"},
	{Name: "src", IsURL: true},
	{Name: "srcdoc"},
	{Name: "srclang"},
	{Name: "srcset"},
	{Name: "start"},
	{Name: "step"},
	{Name: "style"},
	{Name: "tabindex"},
	{Name: "target"},
	{Name: "title"},
	{Name: "translate"},
	{Name: "type"},
	{Name: "usemap", IsURL: true},
	{Name: "value"},
	{Name: "width"},
	{Name: "wrap"},
}

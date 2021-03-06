sleekhtml
=========

Cleans &amp; grooms HTML document from unnecessary white-space, HTML tags, comments &amp; other elements.

[![Build Status](https://travis-ci.org/linkosmos/sleekhtml.svg?branch=master)](https://travis-ci.org/linkosmos/sleekhtml)
[![GoDoc](http://godoc.org/github.com/linkosmos/sleekhtml?status.svg)](http://godoc.org/github.com/linkosmos/sleekhtml)
[![MITLicense](http://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)

## Benchmark (default options)

Website | Raw wget (bytes) | Sanitized (bytes)
------- | ---------------- | ----------------
example.com | 12876 | 531
google.com | 116772 | 2781
economist.com | 194588 | 62278

## Tags
Use tags to get desired output

Defaults:
```go
// DefaultIgnoredHTMLTags -
// Form, input, selects
DefaultIgnoredHTMLTags = []atom.Atom{
	atom.Script, atom.Style, atom.Iframe, atom.Hr,

	atom.Form, atom.Input, atom.Select, atom.Label,
	atom.Fieldset, atom.Button, atom.Textarea,

	atom.Noembed, atom.Embed, atom.Object, atom.Base,
	atom.Canvas, atom.Svg,
}

// DefaultAllowedHTMLAttributes -
// http-equiv, content & charset tags should be always present
// since they handles HTML encoding
DefaultAllowedHTMLAttributes = []string{
	"src", "href", "title", "alt",
	"rel", "http-equiv", "content", "name",
	"description", "charset", "lang",
	"itemprop", "itemscope", "itemref", "itemtype", // Microdata
}

AllowIEComments: false
```

## Usage
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/sleekhtml/sleekhtml"
)

func main() {
	response, err := http.Get("http://www.example.com/")
	defer response.Body.Close()

	tags := sleekhtml.NewTags() // Default Options

	output, err := sleekhtml.Sanitize(response.Body, tags)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
```

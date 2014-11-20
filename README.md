sleekhtml
=========

Cleans &amp; grooms HTML document from unnecessary white-space, HTML tags, comments &amp; other elements.

[![GoDoc](http://godoc.org/github.com/ernestas-poskus/sleekhtml?status.svg)](http://godoc.org/github.com/ernestas-poskus/sleekhtml)
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
IgnoredHTMLTags: []string{
	"script", "style", "iframe", "input",
	"form", "svg", "select", "label",
	"fieldset", "frame", "frameset", "noframes",
	"noembed", "embed", "applet", "object",
	"base",
}

// http-equiv, content & charset tags should be always present
// since they handles HTML encoding
AllowedHTMLAttributes: []string{
	"id", "class", "src", "href",
	"title", "alt", "rel", "http-equiv",
	"content", "name", "description", "charset",
}
```

## Usage
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/ernestas-poskus/sleekhtml"
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

[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/ernestas-poskus/sleekhtml/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

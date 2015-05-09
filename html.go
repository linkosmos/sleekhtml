package sleekhtml

import (
	"bytes"
	"io"
	"strings"

	"code.google.com/p/go.net/html"
)

var emptyByte = []byte("")
var nbsp = []byte("&nbsp;")

// Sanitize - sanitizes & grooms HTML from unnecessary space & tags
func Sanitize(r io.Reader, tags *Tags) ([]byte, error) {
	if tags == nil {
		tags = NewTags()
	}
	return process(html.NewTokenizer(r), tags)
}

func process(tokenizer *html.Tokenizer, tags *Tags) ([]byte, error) {
	var buffer bytes.Buffer
	var ignoredTag = false

	for {
		tt := tokenizer.Next() // TokenType
		if tt == html.ErrorToken {
			return buffer.Bytes(), nil
		}

		switch tt {
		case html.DoctypeToken:
			buffer.Write(trimSpace(tokenizer.Raw()))

		case html.CommentToken: // Ignore HTML comments except IE
			if tags.AllowIEComments && bytes.HasPrefix(tokenizer.Raw(), []byte("<!--[if")) {
				buffer.Write(trimSpace(tokenizer.Raw()))
			}
			break

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if tags.IsIgnoredHTMLTag(token.DataAtom) {
				ignoredTag = true
				break
			}

			if len(token.Attr) > 0 {
				tokenAttr := token.Attr
				token.Attr = nil
				for _, attr := range tokenAttr {
					if tags.IsAllowedAttribute(attr.Key) {
						token.Attr = append(token.Attr, attr)
					}
				}
			}
			buffer.WriteString(strings.TrimSpace(token.String()))
		case html.TextToken:
			if ignoredTag {
				break
			}
			buffer.Write(trimText(tokenizer.Raw()))

		case html.EndTagToken:
			if ignoredTag {
				ignoredTag = false
				break
			}
			buffer.Write(trimSpace(tokenizer.Raw()))
		default:
			buffer.Write(trimSpace(tokenizer.Raw()))
		}
	}
}

func trimSpace(b []byte) []byte {
	return bytes.TrimSpace(b)
}

func trimText(b []byte) []byte {
	return trimSpace(bytes.Replace(b, nbsp, emptyByte, -1))
}

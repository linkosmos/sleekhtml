package sleekhtml

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// FilterTokenFunc - callback function process token
type FilterTokenFunc func(t *html.Token)

var emptyByte = []byte("")
var nbsp = []byte("&nbsp;")

// Sanitize - sanitizes & grooms HTML from unnecessary space & tags
func Sanitize(r io.Reader, tags *Tags) ([]byte, error) {
	if tags == nil {
		tags = NewTags()
	}
	return Process(html.NewTokenizer(r), tags, nil)
}

// Process -
func Process(parser *html.Tokenizer, tags *Tags, tokenFilter FilterTokenFunc) ([]byte, error) {
	var buffer bytes.Buffer
	var ignoredTag = false

	for {
		tt := parser.Next() // TokenType
		if tt == html.ErrorToken {
			return buffer.Bytes(), nil
		}

		switch tt {
		case html.DoctypeToken:
			buffer.Write(trimSpace(parser.Raw()))

		case html.CommentToken: // Ignore HTML comments except IE
			if tags.AllowIEComments && bytes.HasPrefix(parser.Raw(), []byte("<!--[if")) {
				buffer.Write(trimSpace(parser.Raw()))
			}
			break

		case html.StartTagToken, html.SelfClosingTagToken:
			token := parser.Token()
			if tags.IsIgnoredHTMLTag(token.DataAtom) {
				ignoredTag = true
				break
			}

			if tokenFilter != nil {
				tokenFilter(&token)
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
			buffer.Write(trimText(parser.Raw()))

		case html.EndTagToken:
			if ignoredTag {
				ignoredTag = false
				break
			}
			buffer.Write(trimSpace(parser.Raw()))
		default:
			buffer.Write(trimSpace(parser.Raw()))
		}
	}
}

func trimSpace(b []byte) []byte {
	return bytes.TrimSpace(b)
}

func trimText(b []byte) []byte {
	return trimSpace(bytes.Replace(b, nbsp, emptyByte, -1))
}

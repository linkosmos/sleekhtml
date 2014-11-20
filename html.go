package sleekhtml

import (
	"bytes"
	"io"
	"strings"

	"code.google.com/p/go.net/html"
)

// Sanitize - sanitizes & grooms HTML from unnecessary space & tags
func Sanitize(r io.Reader, tags *Tags) (io.Reader, error) {
	if tags == nil {
		tags = NewTags()
	}

	tokenizer := html.NewTokenizer(r)

	var buffer bytes.Buffer
	ignoredTag := false

	for {
		tt := tokenizer.Next() // TokenType
		token := tokenizer.Token()

		switch tt {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return &buffer, nil
			}
			return nil, err

		case html.DoctypeToken:
			buffer.Write(trimSpace(tokenizer.Raw()))

		case html.CommentToken: // Ignore comments (IE as well)
			break

		case html.StartTagToken, html.SelfClosingTagToken:
			if tags.IsIgnoredHTMLTag(token.DataAtom.String()) {
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
			buffer.Write(trimSpace(tokenizer.Raw()))

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

func trimSpace(s []byte) []byte {
	return bytes.TrimSpace(s)
}

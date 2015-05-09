package sleekhtml

import (
	"os"
	"testing"

	"golang.org/x/net/html/atom"
)

func testInput1() (*os.File, string) {
	f, err := os.Open("test-assets/input1.html")
	if err != nil {
		panic(err)
	}
	return f, `<!doctype html><html><head></head><body><div><h1>E ple</h1>AAAA<a href="http://iana.org/do"></p><br/>Hi</div></body></html>`
}

func testTags() *Tags {
	tags := NewTags()
	tags.IgnoredHTMLTags = append(tags.IgnoredHTMLTags, atom.Title)
	tags.IgnoredHTMLTags = append(tags.IgnoredHTMLTags, atom.Meta)
	tags.IgnoredHTMLTags = append(tags.IgnoredHTMLTags, atom.P)

	return tags
}

func TestSanitize(t *testing.T) {
	input, expected := testInput1()
	output, _ := Sanitize(input, testTags())
	got := string(output)

	if got != expected {
		t.Error(len(got), "!=", len(expected))
		t.Error(got, "\n", expected)
	}
}

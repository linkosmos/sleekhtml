package sleekhtml

import (
	"os"
	"testing"
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
	tags.IgnoredHTMLTags = append(tags.IgnoredHTMLTags, "title")
	tags.IgnoredHTMLTags = append(tags.IgnoredHTMLTags, "meta")
	tags.IgnoredHTMLTags = append(tags.IgnoredHTMLTags, "p")

	return tags
}

func TestSanitize(t *testing.T) {
	input, expected := testInput1()
	got, _ := Sanitize(input, testTags())

	if string(got) != expected {
		t.Error(len(got), "!=", len(expected))
		t.Error(got, "\n", expected)
	}
}

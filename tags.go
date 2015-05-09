package sleekhtml

import "golang.org/x/net/html/atom"

var (
	// DefaultIgnoredHTMLTags -
	// Form, input, selects
	// Plugin
	DefaultIgnoredHTMLTags = []atom.Atom{
		atom.Script, atom.Style, atom.Iframe, atom.Hr,

		atom.Form, atom.Input, atom.Select, atom.Label,
		atom.Fieldset, atom.Button,

		atom.Noembed, atom.Embed, atom.Object, atom.Base,
		atom.Canvas, atom.Svg,
	}

	// DefaultAllowedHTMLAttributes -
	// http-equiv, content & charset tags should be always present
	// since they handles HTML encoding
	DefaultAllowedHTMLAttributes = []string{
		"id", "class", "src", "href",
		"title", "alt", "rel", "http-equiv",
		"content", "name", "description", "charset",
		"lang",
		"itemprop", "itemscope", "itemref", "itemtype", // Microdata
	}
)

// Tags - HTML tags structure
type Tags struct {

	// IgnoredHTMLTags - contains tags which will be ignored/removed
	IgnoredHTMLTags []atom.Atom

	// AllowedHTMLAttributes - contains HTML attributes
	AllowedHTMLAttributes []string

	// AllowIEComments - ignore or save IE comments
	AllowIEComments bool
}

// NewTags - initializes Tags with default values
// defaults can be overridden on Tags initialization
func NewTags() *Tags {
	return &Tags{
		IgnoredHTMLTags:       DefaultIgnoredHTMLTags,
		AllowedHTMLAttributes: DefaultAllowedHTMLAttributes,
		AllowIEComments:       false,
	}
}

// IsIgnoredHTMLTag - checks whether tag is not ignored
func (t *Tags) IsIgnoredHTMLTag(a atom.Atom) bool {
	for _, val := range t.IgnoredHTMLTags {
		if val == a {
			return true
		}
	}
	return false
}

// IsAllowedAttribute - checks whether HTML attribute is allowed
func (t *Tags) IsAllowedAttribute(attr string) bool {
	for _, val := range t.AllowedHTMLAttributes {
		if val == attr {
			return true
		}
	}
	return false
}

package sleekhtml

// Tags - HTML tags structure
type Tags struct {

	// IgnoredHTMLTags - contains tags which will be ignored/removed
	IgnoredHTMLTags []string

	// AllowedHTMLAttributes - contains HTML attributes 
	// which will be kept
	AllowedHTMLAttributes []string
}

// NewTags - initializes Tags with default values
// defaults can be overridden on Tags initialization
func NewTags() *Tags {
	return &Tags{
		IgnoredHTMLTags: []string{
			"script", "style", "iframe", "input",
			"form", "svg", "select", "label",
			"fieldset", "frame", "frameset", "noframes",
			"noembed", "embed", "applet", "object",
			"base",
		},
		// http-equiv, content & charset tags should be always present
		// since they handles HTML encoding
		AllowedHTMLAttributes: []string{
			"id", "class", "src", "href",
			"title", "alt", "rel", "http-equiv",
			"content", "name", "description", "charset",
		},
	}
}

// IsIgnoredHTMLTag - checks whether tag is not ignored
func (t *Tags) IsIgnoredHTMLTag(s string) bool {
	return t.includes(t.IgnoredHTMLTags, s)
}

// IsAllowedAttribute - checks whether HTML attribute is allowed
func (t *Tags) IsAllowedAttribute(s string) bool {
	return t.includes(t.AllowedHTMLAttributes, s)
}

func (t *Tags) includes(array []string, element string) bool {
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

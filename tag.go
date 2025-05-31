package defaults

const (
	defaultTag = "default"
)

// Tag specifies the struct tag key used to parse default values for fields.
var Tag = defaultTag

// SetDefaultTag sets the tag name of default tag
func SetDefaultTag(tag string) {
	Tag = tag
}

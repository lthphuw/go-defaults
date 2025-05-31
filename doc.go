// Package defaults provides functionality to parse and set default values for struct fields based on their "default" tags.
// The Defaults function processes a non-nil pointer to a struct, setting default values for exported fields that are unset
// (zero values for non-pointers or nil for pointers) using the tag key specified by the package-level variable Tag (defaulting to "default").
// Nested structs are processed recursively. Fields without a "default" tag are skipped unless they are structs or struct pointers.
//
// Supported field types and example default tags:
//   - int: `default:"123"`
//   - int8: `default:"127"`
//   - int16: `default:"32767"`
//   - int32: `default:"2147483647"`
//   - int64: `default:"9223372036854775807"`
//   - uint: `default:"123"`
//   - uint8: `default:"255"`
//   - uint16: `default:"65535"`
//   - uint32: `default:"4294967295"`
//   - uint64: `default:"18446744073709551615"`
//   - float32: `default:"3.14"`
//   - float64: `default:"3.14159265359"`
//   - complex64: `default:"1+2i"`
//   - complex128: `default:"1.5+2.5i"`
//   - string: `default:"hello"`
//   - bool: `default:"true"`
//   - map (e.g., map[string]any): `default:"{\"key\":\"value\",\"num\":42}"`
//   - slice (e.g., []string): `default:"[\"a\",\"b\",\"c\"]"`
//   - array (e.g., [3]int): `default:"[1,2,3]"`
//   - struct (triggers recursive default setting for nested struct fields)
//   - Pointers to the above types (e.g., *int, *string, *map[string]any, *[]string, *[3]int, *struct):
//   - *int: `default:"123"`
//   - *string: `default:"hello"`
//   - *map[string]any: `default:"{\"key\":\"value\"}"`
//   - *[]string: `default:"[\"a\",\"b\"]"`
//   - *[3]int: `default:"[1,2,3]"`
//   - *struct (recursively processes nested struct fields)
//
// Unsupported field types:
//   - Function types (e.g., func(), *func())
//   - Channels (e.g., chan int)
//   - Interfaces
//   - Unsafe pointers (e.g., unsafe.Pointer)
//   - Any other types not listed above
//
// Default tag values must be valid for the field's type. Numeric types require valid numeric strings, bool requires "true" or "false",
// strings can be plain or JSON-escaped, and maps/slices/arrays require JSON-formatted strings. Errors are returned for invalid inputs,
// unexported fields, empty tags (for non-struct fields), parsing failures (e.g., invalid number formats, JSON syntax errors),
// out-of-range values, or unsupported types. The Defaults function recursively processes nested structs to apply their default tags.
package defaults

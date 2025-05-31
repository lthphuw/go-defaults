package defaults

import (
	"fmt"
	"reflect"
)

// Defaults sets default values for struct fields based on their "default" tags.
//
// It takes a pointer to a struct as input and processes each exported field with a "default" tag.
// If a field is unset (zero value for non-pointers or nil for pointers), the function parses the tag value
// and sets it according to the field's type. Supported types include int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128, bool, string,
// time.Duration, map, slice, array, and nested structs (including pointers to these types).
// The function recursively processes nested structs. It skips unexported fields, non-zero fields,
// and fields without a "default" tag unless they are structs or struct pointers.
//
// Errors are returned for invalid inputs, unsupported types, or parsing failures.
func Defaults(s any) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("input must be a non-nil pointer to a struct")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("input must be a pointer to a struct")
	}

	return setDefaults(v)
}

// setDefaults recursively sets default values for a struct's fields.
func setDefaults(v reflect.Value) error {
	t := v.Type()
	for i := range v.NumField() {
		field := t.Field(i)
		fieldVal := v.Field(i)

		// Skip unexported or unsettable fields
		if !assignable(field, fieldVal) {
			continue
		}

		// Handle nested structs or struct pointers
		if isStructOrStructPtr(fieldVal) &&
			(field.Tag.Get(Tag) == "" || fieldVal.Kind() == reflect.Struct) {
			if fieldVal.Kind() == reflect.Ptr {
				if fieldVal.IsNil() {
					// Initialize nil struct pointer
					fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
				}
				// Recurse into the struct
				if err := setDefaults(fieldVal.Elem()); err != nil {
					return fmt.Errorf("failed to set defaults for field %s: %w", field.Name, err)
				}
			} else {
				// Recurse into the struct
				if err := setDefaults(fieldVal); err != nil {
					return fmt.Errorf("failed to set defaults for field %s: %w", field.Name, err)
				}
			}
			continue
		}

		// Skip if field is not unset (non-zero for non-pointers or non-nil for pointers)
		if !isUnset(fieldVal) {
			continue
		}

		// Get the default tag
		tagVal := field.Tag.Get(Tag)
		if tagVal == "" {
			continue
		}

		// Parse and set the default value
		if err := setFieldValue(fieldVal, field.Type, tagVal); err != nil {
			return fmt.Errorf("failed to set default for field %s: %w", field.Name, err)
		}
	}
	return nil
}

// assignable checks if a field is exported and can be set.
func assignable(field reflect.StructField, fieldVal reflect.Value) bool {
	return field.IsExported() && fieldVal.CanSet()
}

// isUnset checks if a field is unset (nil for pointers, zero value for non-pointers).
func isUnset(val reflect.Value) bool {
	if val.Kind() == reflect.Ptr {
		return val.IsNil()
	}
	return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
}

// isStructOrStructPtr checks if a value is a struct or a pointer to a struct.
func isStructOrStructPtr(val reflect.Value) bool {
	return val.Kind() == reflect.Struct ||
		(val.Kind() == reflect.Ptr && val.Type().Elem().Kind() == reflect.Struct)
}

// setFieldValue parses the tag value and sets it to the field based on its type.
func setFieldValue(fieldVal reflect.Value, fieldType reflect.Type, tagVal string) error {
	// Handle pointer types by initializing and setting the element
	if fieldType.Kind() == reflect.Ptr {
		if fieldVal.IsNil() {
			fieldVal.Set(reflect.New(fieldType.Elem()))
		}
		fieldVal = fieldVal.Elem()
		fieldType = fieldType.Elem()
	}
	var typeName string
	if fieldType.Name() == "" {
		typeName = fieldType.Kind().String()
	} else {
		typeName = fieldType.String()
	}

	// Look up parser function
	parserFunc, exists := parser[typeName]
	if !exists {
		return fmt.Errorf(`%s "%s"`, ErrUnsupportedType.Error(), typeName)
	}

	// Parse the value
	parsedVal, err := parserFunc(tagVal, fieldType)
	if err != nil {
		return err
	}

	// Set the parsed value
	if parsedVal.IsValid() {
		if fieldVal.Kind() == reflect.Map || fieldVal.Kind() == reflect.Slice ||
			fieldVal.Kind() == reflect.Array {
			if parsedVal.Type().ConvertibleTo(fieldType) {
				fieldVal.Set(parsedVal)
			} else {
				return fmt.Errorf("parsed value type %v cannot be converted to field type %v", parsedVal.Type(), fieldType)
			}
		} else {
			fieldVal.Set(parsedVal)
		}
	}

	return nil
}

package defaults

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/segmentio/encoding/json"
)

// ErrUnsupportedType is an error returned when attempting to parse an unsupported type.
var ErrUnsupportedType = errors.New("unsupported type")

// ParserFunc defines a function type for parsing a string into a reflect.Value
// based on the specified reflect.Type, returning the parsed value and any error encountered.
type ParserFunc func(str string, t reflect.Type) (reflect.Value, error)

var parser = map[string]ParserFunc{
	reflect.Int.String():                      ParseInt,
	reflect.Int8.String():                     ParseInt,
	reflect.Int16.String():                    ParseInt,
	reflect.Int32.String():                    ParseInt,
	reflect.Int64.String():                    ParseInt,
	reflect.Uint.String():                     ParseUint,
	reflect.Uint8.String():                    ParseUint,
	reflect.Uint16.String():                   ParseUint,
	reflect.Uint32.String():                   ParseUint,
	reflect.Uint64.String():                   ParseUint,
	reflect.Float32.String():                  ParseFloat,
	reflect.Float64.String():                  ParseFloat,
	reflect.Complex64.String():                ParseComplex,
	reflect.Complex128.String():               ParseComplex,
	reflect.Bool.String():                     ParseBool,
	reflect.String.String():                   ParseString,
	reflect.TypeOf(time.Duration(0)).String(): ParseDuration,
	reflect.Map.String():                      ParseMap,
	reflect.Slice.String():                    ParseSlice,
	reflect.Array.String():                    ParseArray,
}

// ParseInt parses a string to an integer type (int, int8, int16, int32, int64).
func ParseInt(str string, t reflect.Type) (reflect.Value, error) {
	var bitSize int
	switch t.Kind() {
	case reflect.Int8:
		bitSize = 8
	case reflect.Int16:
		bitSize = 16
	case reflect.Int32:
		bitSize = 32
	case reflect.Int64:
		bitSize = 64
	default:
		bitSize = 0
	}
	val, err := strconv.ParseInt(str, 0, bitSize)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(val).Convert(t), nil
}

// ParseUint parses a string to an unsigned integer type (uint, uint8, uint16, uint32, uint64).
func ParseUint(str string, t reflect.Type) (reflect.Value, error) {
	var bitSize int
	switch t.Kind() {
	case reflect.Uint8:
		bitSize = 8
	case reflect.Uint16:
		bitSize = 16
	case reflect.Uint32:
		bitSize = 32
	case reflect.Uint64:
		bitSize = 64
	default:
		bitSize = 0
	}
	val, err := strconv.ParseUint(str, 0, bitSize)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(val).Convert(t), nil
}

// ParseFloat parses a string to a float type (float32, float64).
func ParseFloat(str string, t reflect.Type) (reflect.Value, error) {
	var bitSize int
	switch t.Kind() {
	case reflect.Float32:
		bitSize = 32
	case reflect.Float64:
		bitSize = 64
	default:
		return reflect.Value{}, ErrUnsupportedType
	}
	val, err := strconv.ParseFloat(str, bitSize)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(val).Convert(t), nil
}

// ParseComplex parses a string to a complex type (complex64, complex128).
func ParseComplex(str string, t reflect.Type) (reflect.Value, error) {
	var bitSize int
	switch t.Kind() {
	case reflect.Complex64:
		bitSize = 64
	case reflect.Complex128:
		bitSize = 128
	default:
		return reflect.Value{}, ErrUnsupportedType
	}
	val, err := strconv.ParseComplex(str, bitSize)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(val).Convert(t), nil
}

// ParseBool parses a string to a boolean.
func ParseBool(str string, t reflect.Type) (reflect.Value, error) {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(val).Convert(t), nil
}

// ParseString parses a string to a string.
func ParseString(str string, t reflect.Type) (reflect.Value, error) {
	return reflect.ValueOf(str).Convert(t), nil
}

// ParseDuration parses a string to a time.Duration.
func ParseDuration(str string, t reflect.Type) (reflect.Value, error) {
	val, err := time.ParseDuration(str)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(val).Convert(t), nil
}

// ParseMap parses a JSON-like string to a map.
func ParseMap(str string, t reflect.Type) (reflect.Value, error) {
	if t.Kind() != reflect.Map {
		return reflect.Value{}, fmt.Errorf("t is not a map")
	}
	val := reflect.New(t)
	if err := json.Unmarshal([]byte(str), val.Interface()); err != nil {
		return reflect.Value{}, fmt.Errorf("invalid map format: %w", err)
	}
	return val.Elem(), nil
}

// ParseSlice parses a JSON-like string to a slice.
func ParseSlice(str string, t reflect.Type) (reflect.Value, error) {
	if t.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("t is not a slice")
	}
	val := reflect.New(t)
	if err := json.Unmarshal([]byte(str), val.Interface()); err != nil {
		return reflect.Value{}, fmt.Errorf("invalid slice format: %w", err)
	}
	return val.Elem(), nil
}

// ParseArray parses a JSON-like string to an array.
func ParseArray(str string, t reflect.Type) (reflect.Value, error) {
	if t.Kind() != reflect.Array {
		return reflect.Value{}, fmt.Errorf("t is not an array")
	}
	// Parse into a slice first
	elemType := t.Elem()
	tempSlicePtr := reflect.New(reflect.SliceOf(elemType))
	if err := json.Unmarshal([]byte(str), tempSlicePtr.Interface()); err != nil {
		return reflect.Value{}, fmt.Errorf("invalid array format: %w", err)
	}

	// Check if enough space
	tempSlice := tempSlicePtr.Elem()
	expectedLen := t.Len()
	if tempSlice.Len() > expectedLen {
		return reflect.Value{}, fmt.Errorf(
			"array length %d exceeds capacity %d",
			tempSlice.Len(),
			expectedLen,
		)
	}

	// Append values
	arrayVal := reflect.New(t).Elem()
	for i := range tempSlice.Len() {
		arrayVal.Index(i).Set(tempSlice.Index(i))
	}

	return arrayVal, nil
}

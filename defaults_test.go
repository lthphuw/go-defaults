package defaults

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestSetDefaults(t *testing.T) {
	tests := []struct {
		name      string
		input     *test
		want      *test
		wantErr   bool
		errString string
	}{
		{
			name:  "valid default value",
			input: &test{},
			want: &test{
				Int:        1_000_000,
				Int8:       127,
				Int16:      100,
				Int32:      1_000_000,
				Int64:      100_000_000,
				Uint:       100,
				Uint8:      25,
				Uint16:     123,
				Uint32:     321,
				Uint64:     12_312_312_312,
				Float32:    123.3123,
				Float64:    1234.3123,
				Complex64:  123 + 321i,
				Complex128: 312 + 123i,
				Bool:       true,
				String:     "Hello world",
				Map:        map[string]any{"hi": "hello", "hello": 3.14},
				Slice:      []string{"Hi", "Hi 2", "Hi 3", "Hi 4"},
				Array:      [5]int{1, 2, 3, 4, 5},
				Struct: testStruct{
					A: 20_11_2002,
					B: "gm",
				},
				Duration:      5 * time.Minute,
				IntPtr:        intPtr(1_000_000),
				Int8Ptr:       int8Ptr(127),
				Int16Ptr:      int16Ptr(100),
				Int32Ptr:      int32Ptr(1_000_000),
				Int64Ptr:      int64Ptr(100_000_000),
				UintPtr:       uintPtr(100),
				Uint8Ptr:      uint8Ptr(25),
				Uint16Ptr:     uint16Ptr(123),
				Uint32Ptr:     uint32Ptr(321),
				Uint64Ptr:     uint64Ptr(12_312_312_312),
				Float32Ptr:    float32Ptr(123.3123),
				Float64Ptr:    float64Ptr(1234.3123),
				Complex64Ptr:  complex64Ptr(123 + 321i),
				Complex128Ptr: complex128Ptr(312 + 123i),
				BoolPtr:       boolPtr(true),
				StringPtr:     stringPtr("Hello world"),
				MapPtr:        mapPtr(map[string]any{"hi": "hello", "hello": 3.14}),
				SlicePtr:      slicePtr([]string{"Hi", "Hi 2", "Hi 3", "Hi 4"}),
				ArrayPtr:      arrayPtr([5]int{1, 2, 3, 4, 5}),
				StructPtr: structPtr(testStruct{
					A: 20_11_2002,
					B: "gm",
				}),
				DurationPtr: durationPtr(time.Hour),
			},
			wantErr: false,
		},
		{
			name: "already set value",
			input: &test{
				Int:        1,
				Int8:       1,
				Int16:      1,
				Int32:      1,
				Int64:      2,
				Uint:       3,
				Uint8:      4,
				Uint16:     5,
				Uint32:     6,
				Uint64:     3,
				Float32:    321.3123,
				Float64:    1.3123,
				Complex64:  3 + 321i,
				Complex128: 2 + 123i,
				Bool:       false,
				String:     "Hi",
				Map:        map[string]any{},
				Slice:      []string{"Hi", "Hiii"},
				Array:      [5]int{1, 2, 3, 4, 1000},
				Struct: testStruct{
					A: 123,
					B: "hi",
				},
				Duration:      10 * time.Microsecond,
				IntPtr:        intPtr(1),
				Int8Ptr:       int8Ptr(2),
				Int16Ptr:      int16Ptr(3),
				Int32Ptr:      int32Ptr(4),
				Int64Ptr:      int64Ptr(5),
				UintPtr:       uintPtr(5),
				Uint8Ptr:      uint8Ptr(2),
				Uint16Ptr:     uint16Ptr(3),
				Uint32Ptr:     uint32Ptr(321),
				Uint64Ptr:     uint64Ptr(1),
				Float32Ptr:    float32Ptr(2.3123),
				Float64Ptr:    float64Ptr(21.3123),
				Complex64Ptr:  complex64Ptr(2 + 321i),
				Complex128Ptr: complex128Ptr(1 + 123i),
				BoolPtr:       boolPtr(false),
				StringPtr:     stringPtr("Hello world11"),
				MapPtr:        mapPtr(map[string]any{"222hi": "hello", "hello": 3.14}),
				SlicePtr:      slicePtr([]string{"Hi", "3333Hi 2", "Hi 3", "Hi 4"}),
				ArrayPtr:      arrayPtr([5]int{1, 2323, 3, 4, 5}),
				StructPtr: structPtr(testStruct{
					A: 123,
					B: "hi",
				}),
				DurationPtr: durationPtr(100 * time.Nanosecond),
			},
			want: &test{
				Int:        1,
				Int8:       1,
				Int16:      1,
				Int32:      1,
				Int64:      2,
				Uint:       3,
				Uint8:      4,
				Uint16:     5,
				Uint32:     6,
				Uint64:     3,
				Float32:    321.3123,
				Float64:    1.3123,
				Complex64:  3 + 321i,
				Complex128: 2 + 123i,
				Bool:       true,
				String:     "Hi",
				Map:        map[string]any{},
				Slice:      []string{"Hi", "Hiii"},
				Array:      [5]int{1, 2, 3, 4, 1000},
				Struct: testStruct{
					A: 123,
					B: "hi",
				},
				Duration:      10 * time.Microsecond,
				IntPtr:        intPtr(1),
				Int8Ptr:       int8Ptr(2),
				Int16Ptr:      int16Ptr(3),
				Int32Ptr:      int32Ptr(4),
				Int64Ptr:      int64Ptr(5),
				UintPtr:       uintPtr(5),
				Uint8Ptr:      uint8Ptr(2),
				Uint16Ptr:     uint16Ptr(3),
				Uint32Ptr:     uint32Ptr(321),
				Uint64Ptr:     uint64Ptr(1),
				Float32Ptr:    float32Ptr(2.3123),
				Float64Ptr:    float64Ptr(21.3123),
				Complex64Ptr:  complex64Ptr(2 + 321i),
				Complex128Ptr: complex128Ptr(1 + 123i),
				BoolPtr:       boolPtr(false),
				StringPtr:     stringPtr("Hello world11"),
				MapPtr:        mapPtr(map[string]any{"222hi": "hello", "hello": 3.14}),
				SlicePtr:      slicePtr([]string{"Hi", "3333Hi 2", "Hi 3", "Hi 4"}),
				ArrayPtr:      arrayPtr([5]int{1, 2323, 3, 4, 5}),
				StructPtr: structPtr(testStruct{
					A: 123,
					B: "hi",
				}),
				DurationPtr: durationPtr(100 * time.Nanosecond),
			},
			wantErr: false,
		},
	}

	// Run other tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errString != "" && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("SetDefaults() input = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

// TestNumericOverflowCases tests overflow for all numeric types
func TestNumericOverflowCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "int overflow",
			input:     &testIntOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int8 overflow",
			input:     &testInt8Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int16 overflow",
			input:     &testInt16Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int32 overflow",
			input:     &testInt32Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int64 overflow",
			input:     &testInt64Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint overflow",
			input:     &testUintOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint8 overflow",
			input:     &testUint8Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint16 overflow",
			input:     &testUint16Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint32 overflow",
			input:     &testUint32Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint64 overflow",
			input:     &testUint64Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "float32 overflow",
			input:     &testFloat32Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "float64 overflow",
			input:     &testFloat64Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "complex64 overflow",
			input:     &testComplex64Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "complex128 overflow",
			input:     &testComplex128Overflow{},
			wantErr:   true,
			errString: "value out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestPointerOverflowCases tests overflow for all pointer types
func TestPointerOverflowCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "int pointer overflow",
			input:     &testIntPtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int8 pointer overflow",
			input:     &testInt8PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int16 pointer overflow",
			input:     &testInt16PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int32 pointer overflow",
			input:     &testInt32PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "int64 pointer overflow",
			input:     &testInt64PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint pointer overflow",
			input:     &testUintPtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint8 pointer overflow",
			input:     &testUint8PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint16 pointer overflow",
			input:     &testUint16PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint32 pointer overflow",
			input:     &testUint32PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "uint64 pointer overflow",
			input:     &testUint64PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "float32 pointer overflow",
			input:     &testFloat32PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "float64 pointer overflow",
			input:     &testFloat64PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "complex64 pointer overflow",
			input:     &testComplex64PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
		{
			name:      "complex128 pointer overflow",
			input:     &testComplex128PtrOverflow{},
			wantErr:   true,
			errString: "value out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestInvalidParseCases tests invalid parsing for map and slice fields
func TestInvalidParseCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "invalid map JSON",
			input:     &testInvalidMap{},
			wantErr:   true,
			errString: "invalid character",
		},
		{
			name:      "invalid slice JSON",
			input:     &testInvalidSlice{},
			wantErr:   true,
			errString: "invalid character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestArraySpaceCases tests insufficient array space
func TestArraySpaceCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "array too large",
			input:     &testArrayOverflow{},
			wantErr:   true,
			errString: "array length 6 exceeds capacity 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestPointerCollectionCases tests invalid default values for pointer map, slice, and array
func TestPointerCollectionCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "invalid map pointer",
			input:     &testInvalidMapPtr{},
			wantErr:   true,
			errString: "invalid character",
		},
		{
			name:      "invalid slice pointer",
			input:     &testInvalidSlicePtr{},
			wantErr:   true,
			errString: "invalid character",
		},
		{
			name:      "array pointer overflow",
			input:     &testArrayPtrOverflow{},
			wantErr:   true,
			errString: "exceeds capacity",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestNonStructInputCases tests inputs that are not structs or struct pointers
func TestNonStructInputCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "non-struct input (int)",
			input:     123,
			wantErr:   true,
			errString: "input must be a non-nil pointer to a struct",
		},
		{
			name:      "non-struct input (string)",
			input:     "not a struct",
			wantErr:   true,
			errString: "input must be a non-nil pointer to a struct",
		},
		{
			name:      "non-struct pointer (int pointer)",
			input:     intPtr(123),
			wantErr:   true,
			errString: "input must be a pointer to a struct",
		},
		{
			name:      "not pointer to struct",
			input:     testStruct{},
			wantErr:   true,
			errString: "input must be a non-nil pointer to a struct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestUnexportedFieldCases tests structs with unexported fields
func TestUnexportedFieldCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:  "unexported field",
			input: &testUnexportedField{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestEmptyDefaultTagCases tests structs with empty default tags
func TestEmptyDefaultTagCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:    "empty default tag",
			input:   &testEmptyDefaultTag{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestNestedStructParseError tests parsing error in a nested struct field
func TestNestedStructParseError(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "nested struct parse error",
			input:     &testNestedStructParseError{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "pointer to a struct that parse error",
			input:     &testNestedStructPtrParseError{},
			wantErr:   true,
			errString: "invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

func TestBoolError(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "parse bool error",
			input:     &testParseBool{},
			wantErr:   true,
			errString: "invalid",
		},
		{
			name:      "parse pointer to bool error",
			input:     &testParseBoolPtr{},
			wantErr:   true,
			errString: "invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

// TestFuncTypeCases tests function type and function pointer type with default values
func TestFuncTypeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantErr   bool
		errString string
	}{
		{
			name:      "function type with default",
			input:     &testFuncType{},
			wantErr:   true,
			errString: "unsupported type",
		},
		{
			name:      "function pointer type with default",
			input:     &testFuncPtrType{},
			wantErr:   true,
			errString: "unsupported type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Defaults(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("SetDefaults() error = %v, expected to contain %q", err, tt.errString)
			}
		})
	}
}

func intPtr(x int) *int {
	return &x
}

func int8Ptr(x int8) *int8 {
	return &x
}

func int16Ptr(x int16) *int16 {
	return &x
}

func int32Ptr(x int32) *int32 {
	return &x
}

func int64Ptr(x int64) *int64 {
	return &x
}

func uintPtr(x uint) *uint {
	return &x
}

func uint8Ptr(x uint8) *uint8 {
	return &x
}

func uint16Ptr(x uint16) *uint16 {
	return &x
}

func uint32Ptr(x uint32) *uint32 {
	return &x
}

func uint64Ptr(x uint64) *uint64 {
	return &x
}

func float32Ptr(x float32) *float32 {
	return &x
}

func float64Ptr(x float64) *float64 {
	return &x
}

func complex64Ptr(x complex64) *complex64 {
	return &x
}

func complex128Ptr(x complex128) *complex128 {
	return &x
}

func boolPtr(x bool) *bool {
	return &x
}

func stringPtr(x string) *string {
	return &x
}

func mapPtr(x map[string]any) *map[string]any {
	return &x
}

func slicePtr(x []string) *[]string {
	return &x
}

func arrayPtr(x [5]int) *[5]int {
	return &x
}

func structPtr(x testStruct) *testStruct {
	return &x
}

func durationPtr(x time.Duration) *time.Duration {
	return &x
}

type testStruct struct {
	A int    `default:"20_11_2002"`
	B string `default:"gm"`
}
type test struct {
	Int        int            `default:"1_000_000"`
	Int8       int8           `default:"127"`
	Int16      int16          `default:"100"`
	Int32      int32          `default:"1_000_000"`
	Int64      int64          `default:"100_000_000"`
	Uint       uint           `default:"100"`
	Uint8      uint8          `default:"25"`
	Uint16     uint16         `default:"123"`
	Uint32     uint32         `default:"321"`
	Uint64     uint64         `default:"12_312_312_312"`
	Float32    float32        `default:"123.3123"`
	Float64    float64        `default:"1234.3123"`
	Complex64  complex64      `default:"123+321i"`
	Complex128 complex128     `default:"312+123i"`
	Bool       bool           `default:"true"`
	String     string         `default:"Hello world"`
	Map        map[string]any `default:"{\"hi\": \"hello\", \"hello\": 3.14}"`
	Slice      []string       `default:"[\"Hi\",\"Hi 2\",\"Hi 3\",\"Hi 4\"]"`
	Array      [5]int         `default:"[1, 2, 3, 4, 5]"`
	Duration   time.Duration  `default:"5m"`
	Struct     testStruct

	// Pointers
	IntPtr        *int            `default:"1_000_000"`
	Int8Ptr       *int8           `default:"127"`
	Int16Ptr      *int16          `default:"100"`
	Int32Ptr      *int32          `default:"1_000_000"`
	Int64Ptr      *int64          `default:"100_000_000"`
	UintPtr       *uint           `default:"100"`
	Uint8Ptr      *uint8          `default:"25"`
	Uint16Ptr     *uint16         `default:"123"`
	Uint32Ptr     *uint32         `default:"321"`
	Uint64Ptr     *uint64         `default:"12_312_312_312"`
	Float32Ptr    *float32        `default:"123.3123"`
	Float64Ptr    *float64        `default:"1234.3123"`
	Complex64Ptr  *complex64      `default:"123+321i"`
	Complex128Ptr *complex128     `default:"312+123i"`
	BoolPtr       *bool           `default:"true"`
	StringPtr     *string         `default:"Hello world"`
	MapPtr        *map[string]any `default:"{\"hi\": \"hello\", \"hello\": 3.14}"`
	SlicePtr      *[]string       `default:"[\"Hi\",\"Hi 2\",\"Hi 3\",\"Hi 4\"]"`
	ArrayPtr      *[5]int         `default:"[1, 2, 3, 4, 5]"`
	StructPtr     *testStruct
	DurationPtr   *time.Duration `default:"1h"`
}

// Struct for testing invalid map JSON
type testInvalidMap struct {
	Map map[string]any `default:"{invalid}"`
}

// Struct for testing invalid slice JSON
type testInvalidSlice struct {
	Slice []string `default:"[invalid]"`
}

// Struct for testing array size overflow
type testArrayOverflow struct {
	Array [5]int `default:"[1,2,3,4,5,6]"`
}

// Structs for testing overflow for each numeric type
type testIntOverflow struct {
	Int int `default:"9223372036854775808"` // Exceeds int64 max (2^63-1)
}

type testInt8Overflow struct {
	Int8 int8 `default:"128"` // Exceeds int8 max (127)
}

type testInt16Overflow struct {
	Int16 int16 `default:"32768"` // Exceeds int16 max (32767)
}

type testInt32Overflow struct {
	Int32 int32 `default:"2147483648"` // Exceeds int32 max (2^31-1)
}

type testInt64Overflow struct {
	Int64 int64 `default:"9223372036854775808"` // Exceeds int64 max (2^63-1)
}

type testUintOverflow struct {
	Uint uint `default:"18446744073709551616"` // Exceeds uint64 max (2^64-1)
}

type testUint8Overflow struct {
	Uint8 uint8 `default:"256"` // Exceeds uint8 max (255)
}

type testUint16Overflow struct {
	Uint16 uint16 `default:"65536"` // Exceeds uint16 max (65535)
}

type testUint32Overflow struct {
	Uint32 uint32 `default:"4294967296"` // Exceeds uint32 max (2^32-1)
}

type testUint64Overflow struct {
	Uint64 uint64 `default:"18446744073709551616"` // Exceeds uint64 max (2^64-1)
}

type testFloat32Overflow struct {
	Float32 float32 `default:"3.4e39"` // Exceeds float32 max (~3.4e38)
}

type testFloat64Overflow struct {
	Float64 float64 `default:"1.8e309"` // Exceeds float64 max (~1.8e308)
}

type testComplex64Overflow struct {
	Complex64 complex64 `default:"3.4e39+3.4e39i"` // Exceeds float32 components
}

type testComplex128Overflow struct {
	Complex128 complex128 `default:"1.8e309+1.8e309i"` // Exceeds float64 components
}

// Structs for testing overflow for each pointer type
type testIntPtrOverflow struct {
	IntPtr *int `default:"9223372036854775808"` // Exceeds int64 max (2^63-1)
}

type testInt8PtrOverflow struct {
	Int8Ptr *int8 `default:"128"` // Exceeds int8 max (127)
}

type testInt16PtrOverflow struct {
	Int16Ptr *int16 `default:"32768"` // Exceeds int16 max (32767)
}

type testInt32PtrOverflow struct {
	Int32Ptr *int32 `default:"2147483648"` // Exceeds int32 max (2^31-1)
}

type testInt64PtrOverflow struct {
	Int64Ptr *int64 `default:"9223372036854775808"` // Exceeds int64 max (2^63-1)
}

type testUintPtrOverflow struct {
	UintPtr *uint `default:"18446744073709551616"` // Exceeds uint64 max (2^64-1)
}

type testUint8PtrOverflow struct {
	Uint8Ptr *uint8 `default:"256"` // Exceeds uint8 max (255)
}

type testUint16PtrOverflow struct {
	Uint16Ptr *uint16 `default:"65536"` // Exceeds uint16 max (65535)
}

type testUint32PtrOverflow struct {
	Uint32Ptr *uint32 `default:"4294967296"` // Exceeds uint32 max (2^32-1)
}

type testUint64PtrOverflow struct {
	Uint64Ptr *uint64 `default:"18446744073709551616"` // Exceeds uint64 max (2^64-1)
}

type testFloat32PtrOverflow struct {
	Float32Ptr *float32 `default:"3.4e39"` // Exceeds float32 max (~3.4e38)
}

type testFloat64PtrOverflow struct {
	Float64Ptr *float64 `default:"1.8e309"` // Exceeds float64 max (~1.8e308)
}

type testComplex64PtrOverflow struct {
	Complex64Ptr *complex64 `default:"3.4e39+3.4e39i"` // Exceeds float32 components
}

type testComplex128PtrOverflow struct {
	Complex128Ptr *complex128 `default:"1.8e309+1.8e309i"` // Exceeds float64 components
}

// Struct for testing unexported field
type testUnexportedField struct {
	privateField int `default:"123"`
}

// Struct for testing empty default tag
type testEmptyDefaultTag struct {
	Field int `default:""`
}

// Struct for testing invalid map pointer
type testInvalidMapPtr struct {
	MapPtr *map[string]any `default:"{invalid}"`
}

// Struct for testing invalid slice pointer
type testInvalidSlicePtr struct {
	SlicePtr *[]string `default:"[invalid]"`
}

// Struct for testing array pointer overflow
type testArrayPtrOverflow struct {
	ArrayPtr *[5]int `default:"[1,2,3,4,5,6]"`
}

// Nested struct with invalid default value
type nestedStruct struct {
	A int `default:"invalid_number"` // Invalid numeric value
	B string
}

// Struct containing a nested struct
type testNestedStructParseError struct {
	Nested nestedStruct
}

// Struct containing a nested struct
type testNestedStructPtrParseError struct {
	Nested *nestedStruct
}

// Struct for testing function type with default value
type testFuncType struct {
	FuncField func() `default:"some_value"`
}

// Struct for testing pointer to function type with default value
type testFuncPtrType struct {
	FuncPtrField *func() `default:"some_value"`
}

// Struct for testing boolean error
type testParseBool struct {
	Bool bool `default:"TTT"`
}

// Struct for testing pointer to boolean error
type testParseBoolPtr struct {
	BoolPtr *bool `default:"FAISE"`
}

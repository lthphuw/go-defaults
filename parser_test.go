package defaults

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		typ       reflect.Type
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid small int",
			input:   "42",
			typ:     reflect.TypeOf(int(0)),
			want:    reflect.ValueOf(int(42)),
			wantErr: false,
		},
		{
			name:    "valid small int8",
			input:   "42",
			typ:     reflect.TypeOf(int8(0)),
			want:    reflect.ValueOf(int8(42)),
			wantErr: false,
		},
		{
			name:    "valid small int16",
			input:   "42",
			typ:     reflect.TypeOf(int16(0)),
			want:    reflect.ValueOf(int16(42)),
			wantErr: false,
		},
		{
			name:    "valid small int32",
			input:   "42",
			typ:     reflect.TypeOf(int32(0)),
			want:    reflect.ValueOf(int32(42)),
			wantErr: false,
		},
		{
			name:    "valid small int64",
			input:   "42",
			typ:     reflect.TypeOf(int64(0)),
			want:    reflect.ValueOf(int64(42)),
			wantErr: false,
		},
		{
			name:    "valid negative int",
			input:   "-12",
			typ:     reflect.TypeOf(int(0)),
			want:    reflect.ValueOf(int(-12)),
			wantErr: false,
		},
		{
			name:    "max int8",
			input:   fmt.Sprintf("%d", math.MaxInt8),
			typ:     reflect.TypeOf(int8(0)),
			want:    reflect.ValueOf(int8(math.MaxInt8)),
			wantErr: false,
		},
		{
			name:    "min int8",
			input:   fmt.Sprintf("%d", math.MinInt8),
			typ:     reflect.TypeOf(int8(0)),
			want:    reflect.ValueOf(int8(math.MinInt8)),
			wantErr: false,
		},
		{
			name:    "max int16",
			input:   fmt.Sprintf("%d", math.MaxInt16),
			typ:     reflect.TypeOf(int16(0)),
			want:    reflect.ValueOf(int16(math.MaxInt16)),
			wantErr: false,
		},
		{
			name:    "min int16",
			input:   fmt.Sprintf("%d", math.MinInt16),
			typ:     reflect.TypeOf(int16(0)),
			want:    reflect.ValueOf(int16(math.MinInt16)),
			wantErr: false,
		},
		{
			name:    "max int32",
			input:   fmt.Sprintf("%d", math.MaxInt32),
			typ:     reflect.TypeOf(int32(0)),
			want:    reflect.ValueOf(int32(math.MaxInt32)),
			wantErr: false,
		},
		{
			name:    "min int32",
			input:   fmt.Sprintf("%d", math.MinInt32),
			typ:     reflect.TypeOf(int32(0)),
			want:    reflect.ValueOf(int32(math.MinInt32)),
			wantErr: false,
		},
		{
			name:    "max int64",
			input:   fmt.Sprintf("%d", math.MaxInt64),
			typ:     reflect.TypeOf(int64(0)),
			want:    reflect.ValueOf(int64(math.MaxInt64)),
			wantErr: false,
		},
		{
			name:    "min int64",
			input:   fmt.Sprintf("%d", math.MinInt64),
			typ:     reflect.TypeOf(int64(0)),
			want:    reflect.ValueOf(int64(math.MinInt64)),
			wantErr: false,
		},
		{
			name:      "invalid string input",
			input:     "abc",
			typ:       reflect.TypeOf(int(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "empty string",
			input:     "",
			typ:       reflect.TypeOf(int(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:    "string with _ for int",
			input:   "1_000_000_000",
			typ:     reflect.TypeOf(int(0)),
			want:    reflect.ValueOf(int(1_000_000_000)),
			wantErr: false,
		},
		{
			name:      "positive overflow int8",
			input:     overflowPosInt8(),
			typ:       reflect.TypeOf(int8(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "negative overflow int8",
			input:     overflowNegInt8(),
			typ:       reflect.TypeOf(int8(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow int16",
			input:     overflowPosInt16(),
			typ:       reflect.TypeOf(int16(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "negative overflow int16",
			input:     overflowNegInt16(),
			typ:       reflect.TypeOf(int16(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow int32",
			input:     overflowPosInt32(),
			typ:       reflect.TypeOf(int32(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "negative overflow int32",
			input:     overflowNegInt32(),
			typ:       reflect.TypeOf(int32(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow int64",
			input:     overflowPosInt64(),
			typ:       reflect.TypeOf(int64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "negative overflow int64",
			input:     overflowNegInt64(),
			typ:       reflect.TypeOf(int64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.input, tt.typ)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !errors.Is(err, strconv.ErrSyntax) && tt.errString != "" &&
				!contains(err.Error(), tt.errString) {
				t.Errorf("ParseInt() error = %v, expected to contain %q", err, tt.errString)
			}

			if got.IsValid() != tt.want.IsValid() {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
			if got.IsValid() && tt.want.IsValid() &&
				!reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("ParseInt() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func TestParseUint(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		typ       reflect.Type
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid small uint",
			input:   "42",
			typ:     reflect.TypeOf(uint(0)),
			want:    reflect.ValueOf(uint(42)),
			wantErr: false,
		},
		{
			name:    "valid small uint8",
			input:   "42",
			typ:     reflect.TypeOf(uint8(0)),
			want:    reflect.ValueOf(uint8(42)),
			wantErr: false,
		},
		{
			name:    "valid small uint16",
			input:   "42",
			typ:     reflect.TypeOf(uint16(0)),
			want:    reflect.ValueOf(uint16(42)),
			wantErr: false,
		},
		{
			name:    "valid small uint32",
			input:   "42",
			typ:     reflect.TypeOf(uint32(0)),
			want:    reflect.ValueOf(uint32(42)),
			wantErr: false,
		},
		{
			name:    "valid small uint64",
			input:   "42",
			typ:     reflect.TypeOf(uint64(0)),
			want:    reflect.ValueOf(uint64(42)),
			wantErr: false,
		},
		{
			name:    "zero",
			input:   "0",
			typ:     reflect.TypeOf(uint(0)),
			want:    reflect.ValueOf(uint(0)),
			wantErr: false,
		},
		{
			name:    "max uint8",
			input:   fmt.Sprintf("%d", math.MaxUint8),
			typ:     reflect.TypeOf(uint8(0)),
			want:    reflect.ValueOf(uint8(math.MaxUint8)),
			wantErr: false,
		},
		{
			name:    "max uint16",
			input:   fmt.Sprintf("%d", math.MaxUint16),
			typ:     reflect.TypeOf(uint16(0)),
			want:    reflect.ValueOf(uint16(math.MaxUint16)),
			wantErr: false,
		},
		{
			name:    "max uint32",
			input:   fmt.Sprintf("%d", math.MaxUint32),
			typ:     reflect.TypeOf(uint32(0)),
			want:    reflect.ValueOf(uint32(math.MaxUint32)),
			wantErr: false,
		},
		{
			name:    "max uint64",
			input:   fmt.Sprintf("%d", uint64(math.MaxUint64)),
			typ:     reflect.TypeOf(uint64(0)),
			want:    reflect.ValueOf(uint64(math.MaxUint64)),
			wantErr: false,
		},
		{
			name:      "invalid string input",
			input:     "abc",
			typ:       reflect.TypeOf(uint(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "empty string",
			input:     "",
			typ:       reflect.TypeOf(uint(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "negative number",
			input:     "-1",
			typ:       reflect.TypeOf(uint(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:    "string with _",
			input:   "1_000_000",
			typ:     reflect.TypeOf(uint(0)),
			want:    reflect.ValueOf(uint(1_000_000)),
			wantErr: false,
		},
		{
			name:      "positive overflow uint8",
			input:     overflowPosUint8(),
			typ:       reflect.TypeOf(uint8(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow uint16",
			input:     overflowPosUint16(),
			typ:       reflect.TypeOf(uint16(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow uint32",
			input:     overflowPosUint32(),
			typ:       reflect.TypeOf(uint32(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow uint64",
			input:     overflowPosUint64(),
			typ:       reflect.TypeOf(uint64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUint(tt.input, tt.typ)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !errors.Is(err, strconv.ErrSyntax) && tt.errString != "" &&
				!contains(err.Error(), tt.errString) {
				t.Errorf("ParseUint() error = %v, expected to contain %q", err, tt.errString)
			}

			if got.IsValid() != tt.want.IsValid() {
				t.Errorf("ParseUint() = %v, want %v", got, tt.want)
			}
			if got.IsValid() && tt.want.IsValid() &&
				!reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("ParseUint() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		typ       reflect.Type
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid small float32",
			input:   "42.5",
			typ:     reflect.TypeOf(float32(0)),
			want:    reflect.ValueOf(float32(42.5)),
			wantErr: false,
		},
		{
			name:    "valid small float64",
			input:   "42.5",
			typ:     reflect.TypeOf(float64(0)),
			want:    reflect.ValueOf(float64(42.5)),
			wantErr: false,
		},
		{
			name:    "valid negative float32",
			input:   "-12.3",
			typ:     reflect.TypeOf(float32(0)),
			want:    reflect.ValueOf(float32(-12.3)),
			wantErr: false,
		},
		{
			name:    "zero float64",
			input:   "0.0",
			typ:     reflect.TypeOf(float64(0)),
			want:    reflect.ValueOf(float64(0.0)),
			wantErr: false,
		},
		{
			name:    "max float32",
			input:   fmt.Sprintf("%g", math.MaxFloat32),
			typ:     reflect.TypeOf(float32(0)),
			want:    reflect.ValueOf(float32(math.MaxFloat32)),
			wantErr: false,
		},
		{
			name:    "max float64",
			input:   fmt.Sprintf("%g", math.MaxFloat64),
			typ:     reflect.TypeOf(float64(0)),
			want:    reflect.ValueOf(float64(math.MaxFloat64)),
			wantErr: false,
		},
		{
			name:      "invalid string input",
			input:     "abc",
			typ:       reflect.TypeOf(float32(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "empty string",
			input:     "",
			typ:       reflect.TypeOf(float64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:    "string with _",
			input:   "1_000.5",
			typ:     reflect.TypeOf(float64(0)),
			want:    reflect.ValueOf(float64(1000.5)),
			wantErr: false,
		},
		{
			name:      "positive overflow float32",
			input:     overflowPosFloat32(),
			typ:       reflect.TypeOf(float32(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "positive overflow float64",
			input:     overflowPosFloat64(),
			typ:       reflect.TypeOf(float64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:    "positive infinity float32",
			input:   "Inf",
			typ:     reflect.TypeOf(float32(0)),
			want:    reflect.ValueOf(float32(math.Inf(1))),
			wantErr: false,
		},
		{
			name:    "negative infinity float64",
			input:   "-Inf",
			typ:     reflect.TypeOf(float64(0)),
			want:    reflect.ValueOf(float64(math.Inf(-1))),
			wantErr: false,
		},
		{
			name:    "NaN float32",
			input:   "NaN",
			typ:     reflect.TypeOf(float32(0)),
			wantErr: false,
		},
		{
			name:    "NaN float64",
			input:   "NaN",
			typ:     reflect.TypeOf(float64(0)),
			wantErr: false,
		},
		{
			name:      "unsupported type",
			input:     "42.5",
			typ:       reflect.TypeOf(int(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "unsupported type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFloat(tt.input, tt.typ)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !errors.Is(err, strconv.ErrSyntax) && tt.errString != "" &&
				!contains(err.Error(), tt.errString) {
				t.Errorf("ParseFloat() error = %v, expected to contain %q", err, tt.errString)
			}

			if strings.HasPrefix(tt.name, "NaN") {
				if got.IsValid() && !math.IsNaN(float64(got.Float())) {
					t.Errorf("ParseFloat() = %v, want NaN", got.Interface())
				}
			} else if got.IsValid() && tt.want.IsValid() && !reflect.DeepEqual(got.Interface(), tt.want.Interface()) &&
				!math.IsInf(got.Float(), 0) && !math.IsInf(tt.want.Float(), 0) {
				t.Errorf("ParseFloat() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func TestParseComplex(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		typ       reflect.Type
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid complex64",
			input:   "3+4i",
			typ:     reflect.TypeOf(complex64(0)),
			want:    reflect.ValueOf(complex64(complex(3, 4))),
			wantErr: false,
		},
		{
			name:    "valid complex128",
			input:   "3+4i",
			typ:     reflect.TypeOf(complex128(0)),
			want:    reflect.ValueOf(complex128(complex(3, 4))),
			wantErr: false,
		},
		{
			name:    "valid real only complex64",
			input:   "42",
			typ:     reflect.TypeOf(complex64(0)),
			want:    reflect.ValueOf(complex64(complex(42, 0))),
			wantErr: false,
		},
		{
			name:    "valid imaginary only complex128",
			input:   "5i",
			typ:     reflect.TypeOf(complex128(0)),
			want:    reflect.ValueOf(complex128(complex(0, 5))),
			wantErr: false,
		},
		{
			name:    "negative complex64",
			input:   "-3-4i",
			typ:     reflect.TypeOf(complex64(0)),
			want:    reflect.ValueOf(complex64(complex(-3, -4))),
			wantErr: false,
		},
		{
			name:      "invalid string input",
			input:     "abc",
			typ:       reflect.TypeOf(complex64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "empty string",
			input:     "",
			typ:       reflect.TypeOf(complex128(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:    "string with _",
			input:   "1_000+2_000i",
			typ:     reflect.TypeOf(complex128(0)),
			want:    reflect.ValueOf(complex128(complex(1000, 2000))),
			wantErr: false,
		},
		{
			name:      "overflow real part complex64",
			input:     fmt.Sprintf("%s+1i", overflowPosFloat32()),
			typ:       reflect.TypeOf(complex64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "overflow imaginary part complex64",
			input:     fmt.Sprintf("1+%si", overflowPosFloat32()),
			typ:       reflect.TypeOf(complex64(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:      "overflow real part complex128",
			input:     fmt.Sprintf("%s+1i", overflowPosFloat64()),
			typ:       reflect.TypeOf(complex128(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "out of range",
		},
		{
			name:    "infinity real part complex64",
			input:   "Inf+1i",
			typ:     reflect.TypeOf(complex64(0)),
			want:    reflect.ValueOf(complex64(complex(float32(math.Inf(1)), 1))),
			wantErr: false,
		},
		{
			name:    "NaN real part complex128",
			input:   "NaN+1i",
			typ:     reflect.TypeOf(complex128(0)),
			wantErr: false,
		},
		{
			name:      "unsupported type",
			input:     "3+4i",
			typ:       reflect.TypeOf(int(0)),
			want:      reflect.Value{},
			wantErr:   true,
			errString: "unsupported type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseComplex(tt.input, tt.typ)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseComplex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !errors.Is(err, strconv.ErrSyntax) && tt.errString != "" &&
				!contains(err.Error(), tt.errString) {
				t.Errorf("ParseComplex() error = %v, expected to contain %q", err, tt.errString)
			}

			if strings.HasPrefix(tt.name, "NaN") {
				if got.IsValid() &&
					(!math.IsNaN(float64(real(got.Complex()))) && !math.IsNaN(float64(imag(got.Complex())))) {
					t.Errorf(
						"ParseComplex() = %v, want NaN in real or imaginary part",
						got.Interface(),
					)
				}
			} else if got.IsValid() && tt.want.IsValid() && !reflect.DeepEqual(got.Interface(), tt.want.Interface()) &&
				!math.IsInf(float64(real(got.Complex())), 0) && !math.IsInf(float64(imag(got.Complex())), 0) {
				t.Errorf("ParseComplex() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid true",
			input:   "true",
			want:    reflect.ValueOf(true),
			wantErr: false,
		},
		{
			name:    "valid false",
			input:   "false",
			want:    reflect.ValueOf(false),
			wantErr: false,
		},
		{
			name:    "valid 1",
			input:   "1",
			want:    reflect.ValueOf(true),
			wantErr: false,
		},
		{
			name:    "valid 0",
			input:   "0",
			want:    reflect.ValueOf(false),
			wantErr: false,
		},
		{
			name:    "valid TRUE (case insensitive)",
			input:   "TRUE",
			want:    reflect.ValueOf(true),
			wantErr: false,
		},
		{
			name:    "valid FALSE (case insensitive)",
			input:   "FALSE",
			want:    reflect.ValueOf(false),
			wantErr: false,
		},
		{
			name:      "invalid string input",
			input:     "abc",
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "empty string",
			input:     "",
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
		{
			name:      "invalid number",
			input:     "42",
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBool(tt.input, reflect.TypeOf(false))

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !errors.Is(err, strconv.ErrSyntax) && tt.errString != "" &&
				!contains(err.Error(), tt.errString) {
				t.Errorf("ParseBool() error = %v, expected to contain %q", err, tt.errString)
			}

			if got.IsValid() != tt.want.IsValid() {
				t.Errorf("ParseBool() = %v, want %v", got, tt.want)
			}
			if got.IsValid() && tt.want.IsValid() &&
				!reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("ParseBool() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func TestParseMap(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		targetType reflect.Type
		want       reflect.Value
		wantErr    bool
		errString  string
	}{
		{
			name:       "valid map[string]int",
			input:      `{"a": 1, "b": 2}`,
			targetType: reflect.TypeOf(map[string]int{}),
			want:       reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			wantErr:    false,
		},
		{
			name:       "valid map[string]string",
			input:      `{"key1": "value1", "key2": "value2"}`,
			targetType: reflect.TypeOf(map[string]string{}),
			want:       reflect.ValueOf(map[string]string{"key1": "value1", "key2": "value2"}),
			wantErr:    false,
		},
		{
			name:       "valid empty map",
			input:      `{}`,
			targetType: reflect.TypeOf(map[string]int{}),
			want:       reflect.ValueOf(map[string]int{}),
			wantErr:    false,
		},
		{
			name:       "invalid JSON format",
			input:      `{"a": 1, "b"}`,
			targetType: reflect.TypeOf(map[string]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid map format",
		},
		{
			name:       "empty string",
			input:      "",
			targetType: reflect.TypeOf(map[string]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid map format",
		},
		{
			name:       "non-map target type",
			input:      `{"a": 1}`,
			targetType: reflect.TypeOf([]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "t is not a map",
		},
		{
			name:       "malformed JSON",
			input:      `{a: 1}`,
			targetType: reflect.TypeOf(map[string]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid map format",
		},
		{
			name:       "unquoted string value",
			input:      `{"a": value}`,
			targetType: reflect.TypeOf(map[string]string{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid map format",
		},
		{
			name:       "valid map with mixed types",
			input:      `{"num": 42, "str": "hello"}`,
			targetType: reflect.TypeOf(map[string]interface{}{}),
			want:       reflect.ValueOf(map[string]interface{}{"num": float64(42), "str": "hello"}),
			wantErr:    false,
		},
		{
			name:       "nested map[string]map[string]int",
			input:      `{"outer1": {"inner1": 1, "inner2": 2}, "outer2": {"inner3": 3}}`,
			targetType: reflect.TypeOf(map[string]map[string]int{}),
			want: reflect.ValueOf(map[string]map[string]int{
				"outer1": {"inner1": 1, "inner2": 2},
				"outer2": {"inner3": 3},
			}),
			wantErr: false,
		},
		{
			name:       "map with special characters in keys and values",
			input:      `{"key \"quote\"": "\"special\"", "key,comma": "value_with_underscore"}`,
			targetType: reflect.TypeOf(map[string]string{}),
			want: reflect.ValueOf(map[string]string{
				"key \"quote\"": "\"special\"",
				"key,comma":     "value_with_underscore",
			}),
			wantErr: false,
		},
		{
			name:       "large map with multiple types",
			input:      `{"int": 123, "float": 45.67, "str": "large_map", "bool": true, "null": null}`,
			targetType: reflect.TypeOf(map[string]interface{}{}),
			want: reflect.ValueOf(map[string]interface{}{
				"int":   float64(123),
				"float": float64(45.67),
				"str":   "large_map",
				"bool":  true,
				"null":  nil,
			}),
			wantErr: false,
		},
		{
			name:       "invalid nested map",
			input:      `{"outer": {"inner": invalid}}`,
			targetType: reflect.TypeOf(map[string]map[string]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid map format",
		},
		{
			name:       "complex map with nested mixed types",
			input:      `{"data": {"num": 42, "nested": {"str": "test", "val": 3.14}, "flag": true}}`,
			targetType: reflect.TypeOf(map[string]interface{}{}),
			want: reflect.ValueOf(map[string]interface{}{
				"data": map[string]interface{}{
					"num": float64(42),
					"nested": map[string]interface{}{
						"str": "test",
						"val": float64(3.14),
					},
					"flag": true,
				},
			}),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMap(tt.input, tt.targetType)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errString != "" && !contains(err.Error(), tt.errString) {
				t.Errorf("ParseMap() error = %v, expected to contain %q", err, tt.errString)
			}

			if tt.wantErr {
				if got.IsValid() {
					t.Errorf("ParseMap() = %v, want zero value", got.Interface())
				}
			} else {
				if got.Type() != tt.targetType {
					t.Errorf("ParseMap() returned type = %v, want %v", got.Type(), tt.targetType)
				}
				if !reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
					t.Errorf("ParseMap() = %v, want %v", got.Interface(), tt.want.Interface())
				}
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		targetType reflect.Type
		want       reflect.Value
		wantErr    bool
		errString  string
	}{
		{
			name:       "valid array",
			input:      `[1, 2, 3, 4]`,
			targetType: reflect.TypeOf([4]int{}),
			want:       reflect.ValueOf([4]int{1, 2, 3, 4}),
			wantErr:    false,
		},
		{
			name:       "valid any array",
			input:      `[1, 2, 3.231, "1234"]`,
			targetType: reflect.TypeOf([4]interface{}{}),
			want:       reflect.ValueOf([4]interface{}{float64(1), float64(2), 3.231, "1234"}),
			wantErr:    false,
		},
		{
			name:       "non-array target type",
			input:      `[1, 2, 3.231, "1234"]`,
			targetType: reflect.TypeOf(map[string]string{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "t is not an array",
		},
		{
			name:       "not enough space array",
			input:      `[1, 2, 3, 4, 5]`,
			targetType: reflect.TypeOf([4]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "exceeds capacity",
		},
		{
			name:       "empty array",
			input:      `[]`,
			targetType: reflect.TypeOf([0]string{}),
			want:       reflect.ValueOf([0]string{}),
			wantErr:    false,
		},
		{
			name:       "invalid JSON format",
			input:      `[1, 2,`,
			targetType: reflect.TypeOf([2]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid array format",
		},
		{
			name:       "not a JSON array",
			input:      `{"key": "value"}`,
			targetType: reflect.TypeOf([1]string{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid array format",
		},
		{
			name:       "array of booleans",
			input:      `[true, false]`,
			targetType: reflect.TypeOf([2]bool{}),
			want:       reflect.ValueOf([2]bool{true, false}),
			wantErr:    false,
		},
		{
			name:       "array of floats",
			input:      `[1.1, 2.2, 3.3]`,
			targetType: reflect.TypeOf([3]float64{}),
			want:       reflect.ValueOf([3]float64{1.1, 2.2, 3.3}),
			wantErr:    false,
		},
		{
			name:       "trailing comma",
			input:      `[1, 2, 3,]`,
			targetType: reflect.TypeOf([3]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid array format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArray(tt.input, tt.targetType)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errString != "" && !contains(err.Error(), tt.errString) {
				t.Errorf("ParseArray() error = %v, expected to contain %q", err, tt.errString)
			}

			if tt.wantErr {
				if got.IsValid() {
					t.Errorf("ParseArray() = %v, want zero value", got.Interface())
				}
			} else {
				if got.Type() != tt.targetType {
					t.Errorf("ParseArray() returned type = %v, want %v", got.Type(), tt.targetType)
				}
				if !reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
					t.Errorf("ParseArray() = %v, want %v", got.Interface(), tt.want.Interface())
				}
			}
		})
	}
}

func TestParseSlice(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		targetType reflect.Type
		want       reflect.Value
		wantErr    bool
		errString  string
	}{
		{
			name:       "valid slice",
			input:      `[1, 2, 3, 4]`,
			targetType: reflect.TypeOf([]int{}),
			want:       reflect.ValueOf([]int{1, 2, 3, 4}),
			wantErr:    false,
		},
		{
			name:       "valid any slice",
			input:      `[1, 2, 3.231, "1234"]`,
			targetType: reflect.TypeOf([]interface{}{}),
			want:       reflect.ValueOf([]interface{}{float64(1), float64(2), 3.231, "1234"}),
			wantErr:    false,
		},
		{
			name:       "non-slice target type",
			input:      `[1, 2, 3.231, "1234"]`,
			targetType: reflect.TypeOf([4]interface{}{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "t is not a slice",
		},
		{
			name:       "empty slice",
			input:      `[]`,
			targetType: reflect.TypeOf([]string{}),
			want:       reflect.ValueOf([]string{}),
			wantErr:    false,
		},
		{
			name:       "invalid JSON format",
			input:      `[1, 2,`,
			targetType: reflect.TypeOf([]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid slice format",
		},
		{
			name:       "not a JSON slice",
			input:      `{"key": "value"}`,
			targetType: reflect.TypeOf([]string{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid slice format",
		},
		{
			name:       "slice of booleans",
			input:      `[true, false]`,
			targetType: reflect.TypeOf([]bool{}),
			want:       reflect.ValueOf([]bool{true, false}),
			wantErr:    false,
		},
		{
			name:       "slice of floats",
			input:      `[1.1, 2.2, 3.3]`,
			targetType: reflect.TypeOf([]float64{}),
			want:       reflect.ValueOf([]float64{1.1, 2.2, 3.3}),
			wantErr:    false,
		},
		{
			name:       "trailing comma",
			input:      `[1, 2, 3,]`,
			targetType: reflect.TypeOf([]int{}),
			want:       reflect.Value{},
			wantErr:    true,
			errString:  "invalid slice format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSlice(tt.input, tt.targetType)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errString != "" && !contains(err.Error(), tt.errString) {
				t.Errorf("ParseSlice() error = %v, expected to contain %q", err, tt.errString)
			}

			if tt.wantErr {
				if got.IsValid() {
					t.Errorf("ParseSlice() = %v, want zero value", got.Interface())
				}
			} else {
				if got.Type() != tt.targetType {
					t.Errorf("ParseSlice() returned type = %v, want %v", got.Type(), tt.targetType)
				}
				if !reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
					t.Errorf("ParseSlice() = %v, want %v", got.Interface(), tt.want.Interface())
				}
			}
		})
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid string",
			input:   "hello",
			want:    reflect.ValueOf("hello"),
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			want:    reflect.ValueOf(""),
			wantErr: false,
		},
		{
			name:    "string with spaces",
			input:   "hello world",
			want:    reflect.ValueOf("hello world"),
			wantErr: false,
		},
		{
			name:    "string with special characters",
			input:   `hello\nworld`,
			want:    reflect.ValueOf(`hello\nworld`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseString(tt.input, reflect.TypeOf(""))

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errString != "" && !contains(err.Error(), tt.errString) {
				t.Errorf("ParseString() error = %v, expected to contain %q", err, tt.errString)
			}

			if got.IsValid() != tt.want.IsValid() {
				t.Errorf("ParseString() = %v, want %v", got, tt.want)
			}
			if got.IsValid() && tt.want.IsValid() &&
				!reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("ParseString() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      reflect.Value
		wantErr   bool
		errString string
	}{
		{
			name:    "valid duration",
			input:   "1h30m",
			want:    reflect.ValueOf(1*time.Hour + 30*time.Minute),
			wantErr: false,
		},
		{
			name:    "valid duration with seconds",
			input:   "45s",
			want:    reflect.ValueOf(45 * time.Second),
			wantErr: false,
		},
		{
			name:      "invalid duration",
			input:     "1x",
			want:      reflect.Value{},
			wantErr:   true,
			errString: "unknown unit",
		},
		{
			name:      "empty string",
			input:     "",
			want:      reflect.Value{},
			wantErr:   true,
			errString: "invalid duration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.input, reflect.TypeOf(time.Duration(0)))

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errString != "" && !contains(err.Error(), tt.errString) {
				t.Errorf("ParseDuration() error = %v, expected to contain %q", err, tt.errString)
			}

			if got.IsValid() != tt.want.IsValid() {
				t.Errorf("ParseDuration() = %v, want %v", got, tt.want)
			}
			if got.IsValid() && tt.want.IsValid() &&
				!reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("ParseDuration() = %v, want %v", got.Interface(), tt.want.Interface())
			}
		})
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func overflowPosInt8() string {
	return "128"
}

func overflowNegInt8() string {
	return "-129"
}

func overflowPosInt16() string {
	return "32768"
}

func overflowNegInt16() string {
	return "-32769"
}

func overflowPosInt32() string {
	return "2147483648"
}

func overflowNegInt32() string {
	return "-2147483649"
}

func overflowPosInt64() string {
	return "9223372036854775808"
}

func overflowNegInt64() string {
	return "-9223372036854775809"
}

func overflowPosUint8() string {
	return "256"
}

func overflowPosUint16() string {
	return "65536"
}

func overflowPosUint32() string {
	return "4294967296"
}

func overflowPosUint64() string {
	return "18446744073709551616"
}

func overflowPosFloat32() string {
	return "3.4e39"
}

func overflowPosFloat64() string {
	return "1.8e309"
}

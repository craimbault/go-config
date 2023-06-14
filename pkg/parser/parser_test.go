package parser

import (
	"testing"
)

var StringInvalidValue string = "Invalid Value"

// >--- Bool
var StringToBoolValidValues = []struct {
	a    string
	want bool
}{
	{"true", true},
	{"True", true},
	{"TRUE", true},
	{"1", true},
	{"false", false},
	{"False", false},
	{"FALSE", false},
	{"0", false},
}

func TestStringToBoolValid(t *testing.T) {
	for _, tt := range StringToBoolValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToBool(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %t, want %t", result, tt.want)
			}
		})
	}
}

func TestStringToBoolInvalid(t *testing.T) {
	t.Run(StringInvalidValue, func(t *testing.T) {
		_, err := StringToBool(StringInvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToBoolPanicValid(t *testing.T) {
	StringInvalidValue := "invalid value"
	t.Run(StringInvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToBoolPanic did not panic")
			}
		}()

		StringToBoolPanic(StringInvalidValue)
	})
}

func TestStringToBoolPanicInvalid(t *testing.T) {
	for _, tt := range StringToBoolValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToBoolPanic did panic and is not expected")
				}
			}()

			result := StringToBoolPanic(tt.a)
			if result != tt.want {
				t.Errorf("got %t, want %t", result, tt.want)
			}
		})
	}
}

// <--- Bool

// >--- Float32
var StringToFloat32ValidValues = []struct {
	a    string
	want float32
}{
	{"1.1", 1.1},
	{"1.12345", 1.12345},
	{"1.123456789", 1.123456789},
}

var StringToFloat32InvalidValue string = "Invalid Value"

func TestStringToFloat32Valid(t *testing.T) {
	for _, tt := range StringToFloat32ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToFloat32(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %f, want %f", result, tt.want)
			}
		})
	}
}

func TestStringToFloat32Invalid(t *testing.T) {
	t.Run(StringToFloat32InvalidValue, func(t *testing.T) {
		_, err := StringToFloat32(StringToFloat32InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToFloat32PanicValid(t *testing.T) {
	StringToFloat32InvalidValue := "invalid value"
	t.Run(StringToFloat32InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToFloat32Panic did not panic")
			}
		}()

		StringToFloat32Panic(StringToFloat32InvalidValue)
	})
}

func TestStringToFloat32PanicInvalid(t *testing.T) {
	for _, tt := range StringToFloat32ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToFloat32Panic did panic and is not expected")
				}
			}()

			result := StringToFloat32Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %f, want %f", result, tt.want)
			}
		})
	}
}

// <--- Float32

// >--- Float64
var StringToFloat64ValidValues = []struct {
	a    string
	want float64
}{
	{"1.1", 1.1},
	{"1.12345", 1.12345},
	{"1.123456789", 1.123456789},
}

var StringToFloat64InvalidValue string = "Invalid Value"

func TestStringToFloat64Valid(t *testing.T) {
	for _, tt := range StringToFloat64ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToFloat64(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %f, want %f", result, tt.want)
			}
		})
	}
}

func TestStringToFloat64Invalid(t *testing.T) {
	t.Run(StringToFloat64InvalidValue, func(t *testing.T) {
		_, err := StringToFloat64(StringToFloat64InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToFloat64PanicValid(t *testing.T) {
	StringToFloat64InvalidValue := "invalid value"
	t.Run(StringToFloat64InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToFloat64Panic did not panic")
			}
		}()

		StringToFloat64Panic(StringToFloat64InvalidValue)
	})
}

func TestStringToFloat64PanicInvalid(t *testing.T) {
	for _, tt := range StringToFloat64ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToFloat64Panic did panic and is not expected")
				}
			}()

			result := StringToFloat64Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %f, want %f", result, tt.want)
			}
		})
	}
}

// <--- Float64

// >--- Int
var StringToIntValidValues = []struct {
	a    string
	want int
}{
	{"1", 1},
	{"123", 123},
	{"123456", 123456},
}

var StringToIntInvalidValue string = "Invalid Value"

func TestStringToIntValid(t *testing.T) {
	for _, tt := range StringToIntValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToInt(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToIntInvalid(t *testing.T) {
	t.Run(StringToIntInvalidValue, func(t *testing.T) {
		_, err := StringToInt(StringToIntInvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToIntPanicValid(t *testing.T) {
	StringToIntInvalidValue := "invalid value"
	t.Run(StringToIntInvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToIntPanic did not panic")
			}
		}()

		StringToIntPanic(StringToIntInvalidValue)
	})
}

func TestStringToIntPanicInvalid(t *testing.T) {
	for _, tt := range StringToIntValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToIntPanic did panic and is not expected")
				}
			}()

			result := StringToIntPanic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Int

// >--- Int8
var StringToInt8ValidValues = []struct {
	a    string
	want int8
}{
	{"1", 1},
	{"123", 123},
}

var StringToInt8InvalidValue string = "Invalid Value"

func TestStringToInt8Valid(t *testing.T) {
	for _, tt := range StringToInt8ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToInt8(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToInt8Invalid(t *testing.T) {
	t.Run(StringToInt8InvalidValue, func(t *testing.T) {
		_, err := StringToInt8(StringToInt8InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToInt8PanicValid(t *testing.T) {
	StringToInt8InvalidValue := "invalid value"
	t.Run(StringToInt8InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToInt8Panic did not panic")
			}
		}()

		StringToInt8Panic(StringToInt8InvalidValue)
	})
}

func TestStringToInt8PanicInvalid(t *testing.T) {
	for _, tt := range StringToInt8ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToInt8Panic did panic and is not expected")
				}
			}()

			result := StringToInt8Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Int8

// >--- Int16
var StringToInt16ValidValues = []struct {
	a    string
	want int16
}{
	{"1", 1},
	{"123", 123},
	{"12345", 12345},
}

var StringToInt16InvalidValue string = "Invalid Value"

func TestStringToInt16Valid(t *testing.T) {
	for _, tt := range StringToInt16ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToInt16(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToInt16Invalid(t *testing.T) {
	t.Run(StringToInt16InvalidValue, func(t *testing.T) {
		_, err := StringToInt16(StringToInt16InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToInt16PanicValid(t *testing.T) {
	StringToInt16InvalidValue := "invalid value"
	t.Run(StringToInt16InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToInt16Panic did not panic")
			}
		}()

		StringToInt16Panic(StringToInt16InvalidValue)
	})
}

func TestStringToInt16PanicInvalid(t *testing.T) {
	for _, tt := range StringToInt16ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToInt16Panic did panic and is not expected")
				}
			}()

			result := StringToInt16Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Int16

// >--- Int32
var StringToInt32ValidValues = []struct {
	a    string
	want int32
}{
	{"1", 1},
	{"123", 123},
	{"12345", 12345},
}

var StringToInt32InvalidValue string = "Invalid Value"

func TestStringToInt32Valid(t *testing.T) {
	for _, tt := range StringToInt32ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToInt32(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToInt32Invalid(t *testing.T) {
	t.Run(StringToInt32InvalidValue, func(t *testing.T) {
		_, err := StringToInt32(StringToInt32InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToInt32PanicValid(t *testing.T) {
	StringToInt32InvalidValue := "invalid value"
	t.Run(StringToInt32InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToInt32Panic did not panic")
			}
		}()

		StringToInt32Panic(StringToInt32InvalidValue)
	})
}

func TestStringToInt32PanicInvalid(t *testing.T) {
	for _, tt := range StringToInt32ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToInt32Panic did panic and is not expected")
				}
			}()

			result := StringToInt32Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Int32

// >--- Int64
var StringToInt64ValidValues = []struct {
	a    string
	want int64
}{
	{"1", 1},
	{"123", 123},
	{"12345", 12345},
}

var StringToInt64InvalidValue string = "Invalid Value"

func TestStringToInt64Valid(t *testing.T) {
	for _, tt := range StringToInt64ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToInt64(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToInt64Invalid(t *testing.T) {
	t.Run(StringToInt64InvalidValue, func(t *testing.T) {
		_, err := StringToInt64(StringToInt64InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToInt64PanicValid(t *testing.T) {
	StringToInt64InvalidValue := "invalid value"
	t.Run(StringToInt64InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToInt64Panic did not panic")
			}
		}()

		StringToInt64Panic(StringToInt64InvalidValue)
	})
}

func TestStringToInt64PanicInvalid(t *testing.T) {
	for _, tt := range StringToInt64ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToInt64Panic did panic and is not expected")
				}
			}()

			result := StringToInt64Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Int64

// >--- Uint
var StringToUintValidValues = []struct {
	a    string
	want uint
}{
	{"1", 1},
	{"123", 123},
	{"123456", 123456},
}

var StringToUintInvalidValue string = "Invalid Value"

func TestStringToUintValid(t *testing.T) {
	for _, tt := range StringToUintValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToUint(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToUintInvalid(t *testing.T) {
	t.Run(StringToUintInvalidValue, func(t *testing.T) {
		_, err := StringToUint(StringToUintInvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToUintPanicValid(t *testing.T) {
	StringToUintInvalidValue := "invalid value"
	t.Run(StringToUintInvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToUintPanic did not panic")
			}
		}()

		StringToUintPanic(StringToUintInvalidValue)
	})
}

func TestStringToUintPanicInvalid(t *testing.T) {
	for _, tt := range StringToUintValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToUintPanic did panic and is not expected")
				}
			}()

			result := StringToUintPanic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Uint

// >--- Uint8
var StringToUint8ValidValues = []struct {
	a    string
	want uint8
}{
	{"1", 1},
	{"123", 123},
}

var StringToUint8InvalidValue string = "Invalid Value"

func TestStringToUint8Valid(t *testing.T) {
	for _, tt := range StringToUint8ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToUint8(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToUint8Invalid(t *testing.T) {
	t.Run(StringToUint8InvalidValue, func(t *testing.T) {
		_, err := StringToUint8(StringToUint8InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToUint8PanicValid(t *testing.T) {
	StringToUint8InvalidValue := "invalid value"
	t.Run(StringToUint8InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToUint8Panic did not panic")
			}
		}()

		StringToUint8Panic(StringToUint8InvalidValue)
	})
}

func TestStringToUint8PanicInvalid(t *testing.T) {
	for _, tt := range StringToUint8ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToUint8Panic did panic and is not expected")
				}
			}()

			result := StringToUint8Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Uint8

// >--- Uint16
var StringToUint16ValidValues = []struct {
	a    string
	want uint16
}{
	{"1", 1},
	{"123", 123},
	{"12345", 12345},
}

var StringToUint16InvalidValue string = "Invalid Value"

func TestStringToUint16Valid(t *testing.T) {
	for _, tt := range StringToUint16ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToUint16(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToUint16Invalid(t *testing.T) {
	t.Run(StringToUint16InvalidValue, func(t *testing.T) {
		_, err := StringToUint16(StringToUint16InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToUint16PanicValid(t *testing.T) {
	StringToUint16InvalidValue := "invalid value"
	t.Run(StringToUint16InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToUint16Panic did not panic")
			}
		}()

		StringToUint16Panic(StringToUint16InvalidValue)
	})
}

func TestStringToUint16PanicInvalid(t *testing.T) {
	for _, tt := range StringToUint16ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToUint16Panic did panic and is not expected")
				}
			}()

			result := StringToUint16Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Uint16

// >--- Uint32
var StringToUint32ValidValues = []struct {
	a    string
	want uint32
}{
	{"1", 1},
	{"123", 123},
	{"12345", 12345},
}

var StringToUint32InvalidValue string = "Invalid Value"

func TestStringToUint32Valid(t *testing.T) {
	for _, tt := range StringToUint32ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToUint32(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToUint32Invalid(t *testing.T) {
	t.Run(StringToUint32InvalidValue, func(t *testing.T) {
		_, err := StringToUint32(StringToUint32InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToUint32PanicValid(t *testing.T) {
	StringToUint32InvalidValue := "invalid value"
	t.Run(StringToUint32InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToUint32Panic did not panic")
			}
		}()

		StringToUint32Panic(StringToUint32InvalidValue)
	})
}

func TestStringToUint32PanicInvalid(t *testing.T) {
	for _, tt := range StringToUint32ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToUint32Panic did panic and is not expected")
				}
			}()

			result := StringToUint32Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Uint32

// >--- Uint64
var StringToUint64ValidValues = []struct {
	a    string
	want uint64
}{
	{"1", 1},
	{"123", 123},
	{"12345", 12345},
}

var StringToUint64InvalidValue string = "Invalid Value"

func TestStringToUint64Valid(t *testing.T) {
	for _, tt := range StringToUint64ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			result, err := StringToUint64(tt.a)
			if err != nil {
				t.Errorf("error %s", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

func TestStringToUint64Invalid(t *testing.T) {
	t.Run(StringToUint64InvalidValue, func(t *testing.T) {
		_, err := StringToUint64(StringToUint64InvalidValue)
		if err == nil {
			t.Errorf("no error returned while expected")
		}
	})
}

func TestStringToUint64PanicValid(t *testing.T) {
	StringToUint64InvalidValue := "invalid value"
	t.Run(StringToUint64InvalidValue, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The StringToUint64Panic did not panic")
			}
		}()

		StringToUint64Panic(StringToUint64InvalidValue)
	})
}

func TestStringToUint64PanicInvalid(t *testing.T) {
	for _, tt := range StringToUint64ValidValues {
		t.Run(tt.a, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The StringToUint64Panic did panic and is not expected")
				}
			}()

			result := StringToUint64Panic(tt.a)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

// <--- Uint64

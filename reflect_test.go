package goconfig

import (
	"fmt"
	"reflect"
	"testing"
)

type TestStruct struct {
	Elem1 string `env:"APP_ELEM1"`
}

type TestNestedStruct struct {
	Child TestStruct
	Par1  string `env:"APP_PAR1"`
}

type NotSetableTestStruct struct {
	elem1 string `env:"APP_ELEM1"`
}

func TestGetReflectValueOfValid(t *testing.T) {
	t.Run("TestStruct", func(t *testing.T) {
		test := TestStruct{"elem1"}
		elements := []string{"Elem1"}

		v := getReflectValueOf(&test, elements)
		vTest := reflect.ValueOf(&test).Elem().FieldByName("Elem1")

		if v != vTest {
			t.Errorf("got %+v, want %+v", v, vTest)
		}
	})
}

func TestGetReflectValueOfNestedValid(t *testing.T) {
	t.Run("TestNestedStruct", func(t *testing.T) {
		test := TestNestedStruct{TestStruct{"elem1"}, "par1"}
		elements := []string{"Child", "Elem1"}

		v := getReflectValueOf(&test, elements)
		vTest := reflect.ValueOf(&test).Elem().FieldByName(elements[0]).FieldByName(elements[1])

		if v != vTest {
			t.Errorf("got %+v, want %+v", v, vTest)
		}
	})
}

func TestGetReflectValueOfPanic(t *testing.T) {
	t.Run("TestStruct", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The GetReflectValueOf did not panic")
			}
		}()

		test := NotSetableTestStruct{"elem1"}
		elements := []string{"elem1"}

		getReflectValueOf(&test, elements)
	})
}

func TestSetByKindFromEnvValueValid(t *testing.T) {
	stringValue := "string value"
	floatValue := 1.2
	intValue := 1
	var newValue = any(1)
	v := reflect.ValueOf(&newValue).Elem()
	tests := []struct {
		a    reflect.Kind
		c    string
		want any
	}{
		{reflect.Bool, "true", true},
		{reflect.Float32, fmt.Sprintf("%f", floatValue), float32(floatValue)},
		{reflect.Float64, fmt.Sprintf("%f", floatValue), float64(floatValue)},
		{reflect.Int, fmt.Sprintf("%d", intValue), int(intValue)},
		{reflect.Int8, fmt.Sprintf("%d", intValue), int8(intValue)},
		{reflect.Int16, fmt.Sprintf("%d", intValue), int16(intValue)},
		{reflect.Int32, fmt.Sprintf("%d", intValue), int32(intValue)},
		{reflect.Int64, fmt.Sprintf("%d", intValue), int64(intValue)},
		{reflect.String, stringValue, stringValue},
		{reflect.Uint, fmt.Sprintf("%d", intValue), uint(intValue)},
		{reflect.Uint8, fmt.Sprintf("%d", intValue), uint8(intValue)},
		{reflect.Uint16, fmt.Sprintf("%d", intValue), uint16(intValue)},
		{reflect.Uint32, fmt.Sprintf("%d", intValue), uint32(intValue)},
		{reflect.Uint64, fmt.Sprintf("%d", intValue), uint64(intValue)},
	}

	for _, tt := range tests {
		t.Run(tt.a.String(), func(t *testing.T) {
			err := setByKindFromEnvValue(tt.a, v, tt.c)

			if err != nil {
				t.Errorf("an error occured : %s", err)
			}

			if newValue != tt.want {
				switch tt.a {
				case reflect.Bool, reflect.String:
					t.Errorf("got %s, want %s", newValue, tt.want)
				case reflect.Float32, reflect.Float64:
					t.Errorf("got %f, want %f", newValue, tt.want)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					t.Errorf("got %d, want %d", newValue, tt.want)
				}
			}
		})
	}
}

func TestSetByKindFromEnvValueError(t *testing.T) {
	var newValue = any(1)
	v := reflect.ValueOf(&newValue).Elem()

	t.Run("Invalid", func(t *testing.T) {
		err := setByKindFromEnvValue(reflect.Array, v, "TEST")

		if err == nil {
			t.Errorf("no error returned")
		}
	})
}

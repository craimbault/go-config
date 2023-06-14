package goconfig

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/ini.v1"
)

type TestStruct struct {
	Elem1 string `env:"APP_ELEM1" ini:"test>elem1"`
}

type TestStruct2 struct {
	Elem1 string `env:"APP_ELEM1" ini:"test>elem1"`
	Elem2 string `env:"APP_ELEM2" ini:"elem2"`
}

type TestStructFieldTypeUnkown struct {
	Elem1 any `env:"APP_ELEM1" ini:"test>elem1"`
}

type TestNestedStruct struct {
	Par1  string `env:"APP_PAR1" ini:"par1"`
	Child TestStruct
}

type NoEnvNoIniTestStruct struct {
	Par1 string `test:"test"`
}

type NotSetableTestStruct struct {
	elem1 string `env:"APP_ELEM1"`
}

var testedValues []string

func replaceWithValueTest(e interface{}, field reflect.StructField, parents []string, override interface{}) {
	testedValues = append(testedValues, strings.Join(append(parents, field.Name), "."))
}

func TestGetReflectValueOf(t *testing.T) {
	t.Run("SimpleStruct", func(t *testing.T) {
		test := TestStruct{"elem1"}
		elements := []string{"Elem1"}

		v := getReflectValueOf(&test, elements)
		vTest := reflect.ValueOf(&test).Elem().FieldByName("Elem1")

		if v != vTest {
			t.Errorf("got %+v, want %+v", v, vTest)
		}
	})
	t.Run("NestedStruct", func(t *testing.T) {
		test := TestNestedStruct{"par1", TestStruct{"elem1"}}
		elements := []string{"Child", "Elem1"}

		v := getReflectValueOf(&test, elements)
		vTest := reflect.ValueOf(&test).Elem().FieldByName(elements[0]).FieldByName(elements[1])

		if v != vTest {
			t.Errorf("got %+v, want %+v", v, vTest)
		}
	})
	t.Run("Panic", func(t *testing.T) {
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

func TestSetByKindFromStringValueValid(t *testing.T) {
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
			err := setByKindFromStringValue(tt.a, v, tt.c)

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

	newValue = any(1)
	v = reflect.ValueOf(&newValue).Elem()

	t.Run("Invalid", func(t *testing.T) {
		err := setByKindFromStringValue(reflect.Array, v, "TEST")

		if err == nil {
			t.Errorf("no error returned")
		}
	})
}

func TestReflectStructWalk(t *testing.T) {
	t.Run("TestPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The ReflectStructWalkError did not panic")
			}
		}()

		test := "NotAStruct"
		var parents []string

		reflectStructWalk(
			&test,
			reflect.TypeOf(&test).Elem(),
			replaceWithValueTest,
			parents,
			nil,
		)
	})
	t.Run("TestIsValid", func(t *testing.T) {

		var (
			parents []string
			config  = TestNestedStruct{}
		)

		// On reset les resultats
		testedValues = nil

		// On definit celles attendues
		expectedValues := []string{
			"Child.Elem1",
			"Par1",
		}

		// On lance le test
		reflectStructWalk(
			&config,
			reflect.TypeOf(&config).Elem(),
			replaceWithValueTest,
			parents,
			nil,
		)

		// Si l'on a pas le bon nombre de resultats
		if len(testedValues) != len(expectedValues) {
			t.Error("Not enough results in testedValues")
		}

		// On parcours les valeurs attendues
		for _, expected := range expectedValues {
			present := false
			// On parcours les valeurs generees
			for _, available := range testedValues {
				if expected == available {
					present = true
				}
			}
			// Si l'on a pas trouve l'element
			if !present {
				t.Errorf("Missing tested value : %s", expected)
			}
		}
	})
}

func TestReplaceWithEnvValue(t *testing.T) {
	t.Run("TestPanicNoEnvDefined", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The replaceWithEnvValue did not panic")
			}
		}()

		var (
			parents []string
			config  = NoEnvNoIniTestStruct{}
		)

		replaceWithEnvValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			nil,
		)
	})

	t.Run("TestEnvNotDefined", func(t *testing.T) {
		var (
			parents []string
			config  = TestStruct{
				Elem1: "TEST",
			}
		)

		replaceWithEnvValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			nil,
		)

		if config.Elem1 != "TEST" {
			t.Errorf("Value should not be modified")
		}
	})

	t.Run("TestEnvDefined", func(t *testing.T) {
		var (
			parents []string
			config  = TestStruct{
				Elem1: "TEST",
			}
		)

		// On indique l'env
		os.Setenv("APP_ELEM1", "NEW")

		replaceWithEnvValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			nil,
		)

		if config.Elem1 != "NEW" {
			t.Errorf("Value should have been modified")
		}
	})

	t.Run("TestEnvDefinedUnknownType", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The replaceWithEnvValue did not panic")
			}
		}()

		var (
			parents []string
			config  = TestStructFieldTypeUnkown{}
		)

		replaceWithEnvValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			nil,
		)
	})
}

func TestReplaceWithIniValue(t *testing.T) {
	validIni := ini.Empty()
	validIni.Section("").Key("elem2").SetValue("NEW2")
	section, _ := validIni.NewSection("test")
	section.Key("elem1").SetValue("NEW1")
	invalidIni := ini.Empty()

	t.Run("TestPanicNoIniDefined", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The replaceWithIniValue did not panic")
			}
		}()

		var (
			parents []string
			config  = NoEnvNoIniTestStruct{}
		)

		replaceWithIniValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			validIni,
		)
	})

	t.Run("TestSectionNotDefined", func(t *testing.T) {
		var (
			parents []string
			config  = TestStruct2{
				Elem1: "TEST",
				Elem2: "TEST",
			}
		)

		replaceWithIniValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			invalidIni,
		)

		if config.Elem1 != "TEST" || config.Elem2 != "TEST" {
			t.Errorf("Values should not be modified")
		}
	})

	t.Run("TestAllDefined", func(t *testing.T) {
		var (
			parents []string
			config  = TestStruct2{
				Elem1: "TEST",
				Elem2: "TEST",
			}
		)

		replaceWithIniValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			validIni,
		)

		if config.Elem1 != "NEW1" {
			t.Errorf("Elem1 has not been modified")
		}

		replaceWithIniValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(1),
			parents,
			validIni,
		)

		if config.Elem2 != "NEW2" {
			t.Errorf("Elem2 has not been modified")
		}
	})

	t.Run("TestEnvDefinedUnknownType", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The replaceWithIniValue did not panic")
			}
		}()

		var (
			parents []string
			config  = TestStructFieldTypeUnkown{}
		)

		replaceWithIniValue(
			&config,
			reflect.TypeOf(&config).Elem().Field(0),
			parents,
			validIni,
		)
	})
}

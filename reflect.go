package goconfig

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/craimbault/go-config/parser"
)

func reflectStructWalk(e any, t reflect.Type, extractMethod structExtract, parents []string) {
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Type.Kind() == reflect.Struct {
				parents = append(parents, f.Name)
				reflectStructWalk(e, f.Type, extractMethod, parents)
			} else {
				if f.Tag != "" {
					extractMethod(e, f, parents)
				}
			}
		}
	} else {
		panic(errors.New("[go-config] struct walk element is not a struct"))
	}
}

type structExtract func(interface{}, reflect.StructField, []string)

func replaceWithEnvValue(e interface{}, field reflect.StructField, parents []string) {
	// On recupere la valeur
	envName, exists := field.Tag.Lookup("env")

	if !exists {
		panic(errors.New("[go-config] struct tag 'env' is not defined in field[" + field.Name + "]"))
	}

	// On recupere la valeur depuis l'ENV
	envStrVal := os.Getenv(envName)

	// Si l'on a une info depuis l'ENV
	if len(envStrVal) > 0 {
		// On ajoute le champ courant a la liste des champs a parcourir pour faire le set (recursif)
		fieldNames := append(parents, field.Name)

		// On recupere la reflection
		v := getReflectValueOf(e, fieldNames)

		// On change la valeur
		err := setByKindFromEnvValue(field.Type.Kind(), v, envStrVal)
		if err != nil {
			panic(errors.New("[go-config] Type[" + field.Type.Name() + "] not handled for field[" + field.Name + "]"))
		}
	}
}

func getReflectValueOf(e interface{}, fieldNames []string) reflect.Value {
	v := reflect.ValueOf(e).Elem()
	for _, fieldName := range fieldNames {
		v = v.FieldByName(fieldName)
	}

	// Si l'on a pas access en ecriture
	if !v.CanSet() {
		panic(errors.New("[go-config] cannot set field in [" + strings.Join(fieldNames, ".") + "]"))
	}

	return v
}

func setByKindFromEnvValue(k reflect.Kind, v reflect.Value, s string) error {
	var err error = nil

	// On parse la valeur en fonction du type
	switch k {
	case reflect.Bool:
		v.Set(reflect.ValueOf(parser.StringToBoolPanic(s)))
	case reflect.Float32:
		v.Set(reflect.ValueOf(parser.StringToFloat32Panic(s)))
	case reflect.Float64:
		v.Set(reflect.ValueOf(parser.StringToFloat64Panic(s)))
	case reflect.Int:
		v.Set(reflect.ValueOf(parser.StringToIntPanic(s)))
	case reflect.Int8:
		v.Set(reflect.ValueOf(parser.StringToInt8Panic(s)))
	case reflect.Int16:
		v.Set(reflect.ValueOf(parser.StringToInt16Panic(s)))
	case reflect.Int32:
		v.Set(reflect.ValueOf(parser.StringToInt32Panic(s)))
	case reflect.Int64:
		v.Set(reflect.ValueOf(parser.StringToInt64Panic(s)))
	case reflect.String:
		v.Set(reflect.ValueOf(s))
	case reflect.Uint:
		v.Set(reflect.ValueOf(parser.StringToUintPanic(s)))
	case reflect.Uint8:
		v.Set(reflect.ValueOf(parser.StringToUint8Panic(s)))
	case reflect.Uint16:
		v.Set(reflect.ValueOf(parser.StringToUint16Panic(s)))
	case reflect.Uint32:
		v.Set(reflect.ValueOf(parser.StringToUint32Panic(s)))
	case reflect.Uint64:
		v.Set(reflect.ValueOf(parser.StringToUint64Panic(s)))
	default:
		err = errors.New("[go-config] Type[" + k.String() + "] not handled")
	}

	return err
}

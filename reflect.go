package goconfig

import (
	"errors"
	"os"
	"reflect"

	"github.com/craimbault/go-config/parser"
)

func reflectWalk(e interface{}, t reflect.Type) {
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		reflectWalk(e, t.Elem())
	case reflect.Struct:
		var parents []string
		reflectStructWalk(e, t, replaceWithEnvValue, parents)
	}
}

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
		v := reflect.ValueOf(e).Elem()
		for _, fieldName := range fieldNames {
			v = v.FieldByName(fieldName)
		}

		// Si l'on a pas access en ecriture
		if !v.CanSet() {
			panic(errors.New("[go-config] cannot set field[" + field.Name + "]"))
		}

		// On parse la valeur en fonction du type
		switch field.Type.Kind() {
		case reflect.Bool:
			v.SetBool(parser.StringToBoolPanic(envStrVal))
		case reflect.Float32:
			v.Set(reflect.ValueOf(parser.StringToFloat32Panic(envStrVal)))
		case reflect.Float64:
			v.Set(reflect.ValueOf(parser.StringToFloat64Panic(envStrVal)))
		case reflect.Int:
			v.Set(reflect.ValueOf(parser.StringToIntPanic(envStrVal)))
		case reflect.Int8:
			v.Set(reflect.ValueOf(parser.StringToInt8Panic(envStrVal)))
		case reflect.Int16:
			v.Set(reflect.ValueOf(parser.StringToInt16Panic(envStrVal)))
		case reflect.Int32:
			v.Set(reflect.ValueOf(parser.StringToInt32Panic(envStrVal)))
		case reflect.Int64:
			v.Set(reflect.ValueOf(parser.StringToInt64Panic(envStrVal)))
		case reflect.String:
			v.SetString(envStrVal)
		case reflect.Uint:
			v.Set(reflect.ValueOf(parser.StringToUintPanic(envStrVal)))
		case reflect.Uint8:
			v.Set(reflect.ValueOf(parser.StringToUint8Panic(envStrVal)))
		case reflect.Uint16:
			v.Set(reflect.ValueOf(parser.StringToUint16Panic(envStrVal)))
		case reflect.Uint32:
			v.Set(reflect.ValueOf(parser.StringToUint32Panic(envStrVal)))
		case reflect.Uint64:
			v.Set(reflect.ValueOf(parser.StringToUint64Panic(envStrVal)))
		default:
			panic(errors.New("[go-config] Type[" + field.Type.Name() + "] not handled for field[" + field.Name + "]"))
		}
	}
}

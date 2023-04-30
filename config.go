package goconfig

import (
	"reflect"

	"github.com/mcuadros/go-defaults"
)

func LoadConfig(config interface{}) {
	// On charge les valeurs par defaut
	defaults.SetDefaults(config)

	// On ecrase avec les variables d'ENV si definies
	var parents []string
	reflectStructWalk(config, reflect.TypeOf(config).Elem(), replaceWithEnvValue, parents)
}

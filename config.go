package goconfig

import (
	"reflect"

	"github.com/mcuadros/go-defaults"
)

func LoadConfig(config interface{}) {
	// On charge les valeurs par defaut
	defaults.SetDefaults(config)

	// On regarde dans les variables d'env
	reflectWalk(config, reflect.TypeOf(config))
}

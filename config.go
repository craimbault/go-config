package goconfig

import (
	"reflect"

	"github.com/mcuadros/go-defaults"
	"gopkg.in/ini.v1"
)

func LoadConfig(config interface{}, useIniFiles []string, useEnv bool) {
	var parents []string

	// On charge les valeurs par defaut
	defaults.SetDefaults(config)

	// On passe dans les ini a charger
	for _, iniFile := range useIniFiles {
		// On reset les parents
		parents = nil

		// On charge la config
		cfg, err := ini.Load(iniFile)

		// Si l'on a pas d'erreur sur la config
		if err == nil {
			// On ecrase avec les variables d'ENV si definies
			reflectStructWalk(
				config,
				reflect.TypeOf(config).Elem(),
				replaceWithIniValue,
				parents,
				cfg,
			)
		}
	}

	// Si l'on remplace avec les env
	if useEnv {
		parents = nil
		// On ecrase avec les variables d'ENV si definies
		reflectStructWalk(
			config,
			reflect.TypeOf(config).Elem(),
			replaceWithEnvValue,
			parents,
			nil,
		)
	}
}

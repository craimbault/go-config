package goconfig

import (
	"os"
	"testing"
)

type AppConfig struct {
	Name string `env:"APP_NAME" ini:"name" default:"MyApp"`
}

func TestLoadConfig(t *testing.T) {
	t.Run("WithDefault", func(t *testing.T) {
		var config = AppConfig{}
		iniFiles := []string{}

		LoadConfig(&config, iniFiles, false)

		if config.Name != "MyApp" {
			t.Error("Default not valid")
		}
	})

	t.Run("WithIniOverload", func(t *testing.T) {
		var config = AppConfig{}
		iniFiles := []string{
			"testdata/test.ini",
		}

		LoadConfig(&config, iniFiles, false)

		if config.Name != "test" {
			t.Error("Ini file not loaded")
		}
	})

	t.Run("WithMultipleIniOverload", func(t *testing.T) {
		var config = AppConfig{}
		iniFiles := []string{
			"testdata/test.ini",
			"testdata/test2.ini",
		}

		LoadConfig(&config, iniFiles, false)

		if config.Name != "test2" {
			t.Error("Ini file not loaded in the right order")
		}
	})

	t.Run("WithIniThenEnvOverload", func(t *testing.T) {
		var config = AppConfig{}
		iniFiles := []string{
			"testdata/test.ini",
		}

		os.Setenv("APP_NAME", "env")
		LoadConfig(&config, iniFiles, true)

		if config.Name != "env" {
			t.Error("Env is not the last config loaded")
		}
	})
}

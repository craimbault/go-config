package main

import (
	"fmt"

	goconfig "github.com/craimbault/go-config"
)

type AppConfig struct {
	Name        string `env:"APP_NAME" ini:"name" default:"MyApp"`
	Environment string `env:"APP_ENV" ini:"env" default:"local"`
	Debug       bool   `env:"APP_DEBUG" ini:"debug" default:"true"`
	Url         string `env:"APP_URL"  ini:"url" default:"http://localhost"`
	Port        int64  `env:"APP_PORT" ini:"port" default:"8098"`
	TimeZone    string `env:"APP_TIMEZONE" ini:"tz" default:"Europe/Paris"`
	Locale      string `env:"APP_LOCALE" ini:"locale" default:"fr"`
	GinMode     string `env:"GIN_MODE" ini:"gin_mode" default:"debug"`
	Mongo       MongoConfig
}

type MongoConfig struct {
	Uri      string `env:"APP_MONGO_URI" ini:"mongo>uri" default:"mongodb://admuser:admpwd@documentdb:27017/?authSource=admin"`
	Database string `env:"APP_MONGO_DATABASE" ini:"mongo>database" default:"example"`
}

func main() {
	var config = AppConfig{}
	iniFiles := []string{
		"example.ini",
	}

	goconfig.LoadConfig(&config, iniFiles, true)

	fmt.Println(config)
}

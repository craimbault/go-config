Enabling Config by stuctures with defaults values, read from ini files and read from ENV using [struct tags](http://golang.org/pkg/reflect/#StructTag).

> INI file Usage :
>
>> Access INI Section then Key by using `section>key`.
>>
>> If Key has no section, just provide Key name and value.

Installation
------------

The recommended way to install go-config

```bash
go get github.com/craimbault/go-config
```

Examples
--------

A basic example:

```go

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
		"main.ini",
	}

	goconfig.LoadConfig(&config, iniFiles, true)

	fmt.Println(config)
}

```

TODO
--------


License
-------

MIT, see [LICENSE](LICENSE)
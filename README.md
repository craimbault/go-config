Enabling Config by stuctures with defaults values and read from ENV using [struct tags](http://golang.org/pkg/reflect/#StructTag).

Installation
------------

The recommended way to install go-config

```
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
	Name        string `env:"APP_NAME" default:"MyApp"`
	Environment string `env:"APP_ENV" default:"local"`
	Debug       bool   `env:"APP_DEBUG" default:"true"`
	Url         string `env:"APP_URL"  default:"http://localhost"`
	Port        int64  `env:"APP_PORT" default:"8098"`
	TimeZone    string `env:"APP_TIMEZONE"  default:"Europe/Paris"`
	Locale      string `env:"APP_LOCALE"  default:"fr"`
	GinMode     string `env:"GIN_MODE" default:"debug"`
	Mongo       MongoConfig
}

type MongoConfig struct {
	Uri      string `env:"APP_MONGO_URI" default:"mongodb://admuser:admpwd@documentdb:27017/?authSource=admin"`
	Database string `env:"APP_MONGO_DATABASE" default:"example"`
}

func main() {
	var config = AppConfig{}

	goconfig.LoadConfig(&config)

	fmt.Println(config)
}
```

TODO
--------
- Add tests

License
-------

MIT, see [LICENSE](LICENSE)
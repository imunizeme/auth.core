package config

import (
	"strings"

	"github.com/crgimenes/goConfig"
	// toml import for goConfig
	_ "github.com/crgimenes/goConfig/toml"
	"github.com/nuveo/log"
	pConf "github.com/prest/config"
)

// Config for API
type Config struct {
	Cors     string `toml:"cors" cfg:"cors" cfgDefault:"*"`
	Debug    bool   `toml:"debug" cfg:"debug" cfgDefault:"false"`
	JWTKey   string `toml:"jwt_key" cfg:"jwt_key"`
	PGHost   string `toml:"pg_host" cfg:"pg_host" cfgDefault:"127.0.0.1"`
	PGPort   int    `toml:"pg_port" cfg:"pg_port" cfgDefault:"5432"`
	PGDBName string `toml:"pg_dbname" cfg:"pg_dbname"`
	PGUser   string `toml:"pg_user" cfg:"pg_user" cfgDefault:"postgres"`
	Prest    Prest  `toml:"prest" cfg:"prest"`
	Auth     Auth   `toml:"auth" cfg:"auth"`
}

// Prest config
type Prest struct {
	Host       string `toml:"host" cfg:"host" cfgDefault:"127.0.0.1"`
	Port       int    `toml:"port" cfg:"port" cfgDefault:"3000"`
	Migrations string `toml:"migrations" cfg:"migrations" cfgDefault:"./migrations"`
	Queries    string `toml:"queries" cfg:"queries" cfgDefault:"./queries"`
}

// Auth config
type Auth struct {
	Host string `toml:"host" cfg:"host" cfgDefault:"127.0.0.1"`
	Port int    `toml:"port" cfg:"port" cfgDefault:"4000"`
}

// Get cconfig global var
var Get *Config

// Load configs
func Load() (err error) {
	goConfig.PrefixEnv = "IMUNIZEME"
	goConfig.File = "config.toml"
	Get = &Config{}

	err = goConfig.Parse(Get)
	if err != nil {
		return
	}

	if Get.Debug {
		log.DebugMode = true
		log.Println("DEBUG MODE ON")
	}

	pConf.Load()
	pConf.PrestConf.HTTPPort = Get.Prest.Port
	pConf.PrestConf.PGHost = Get.PGHost
	pConf.PrestConf.PGPort = Get.PGPort
	pConf.PrestConf.PGDatabase = Get.PGDBName
	pConf.PrestConf.PGUser = Get.PGUser
	pConf.PrestConf.JWTKey = Get.JWTKey
	pConf.PrestConf.Debug = Get.Debug
	pConf.PrestConf.CORSAllowOrigin = strings.Fields(Get.Cors)
	pConf.PrestConf.MigrationsPath = Get.Prest.Migrations
	pConf.PrestConf.QueriesPath = Get.Prest.Queries
	return
}

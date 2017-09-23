package config

import (
	"strings"

	"github.com/crgimenes/goconfig"
	// toml import for goConfig
	_ "github.com/crgimenes/goconfig/toml"
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
	Bot      Bot    `toml:"bot" cfg:"bot"`
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

// Bot config
type Bot struct {
	Consumerkey    string `toml:"consumer_key" cfg:"consumer_key"`
	ConsumerSecret string `toml:"consumer_secret" cfg:"consumer_secret"`
	AccessToken    string `toml:"access_token" cfg:"access_token"`
	TokenSecret    string `toml:"token_secret" cfg:"token_secret"`
	MessageToken   string `toml:"message_token" cfg:"message_token"`
}

// Get cconfig global var
var Get *Config

// Load configs
func Load() (err error) {
	goconfig.PrefixEnv = "IMUNIZEME"
	goconfig.File = "config.toml"
	Get = &Config{}

	err = goconfig.Parse(Get)
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

package util

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var configName = "devConfig"

func init() {
	flag.StringVar(&configName, "devConfig", "devConfig", "config file name without extension")
	flag.Parse()
}

type Cfg struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	PostgresDB    string `mapstructure:"POSTGRES_DB"`
	PostgresUser  string `mapstructure:"POSTGRES_USER"`
	PostgresPass  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost  string `mapstructure:"POSTGRES_HOST"`
	PostgresPort  string `mapstructure:"POSTGRES_PORT"`
}

var Config Cfg

func LoadConfig() {
	viper.SetConfigName(configName)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("fatal error config file: ", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal("unable to decode into struct: ", err)
	}
}

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		Config.PostgresUser,
		Config.PostgresPass,
		Config.PostgresHost,
		Config.PostgresPort,
		Config.PostgresDB,
	)
}

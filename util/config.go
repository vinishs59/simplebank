package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver	string  `mapstructure:"DB_DRIVER"`
	DBSource	string	`mapstructure:"DB_SOURCE"`
	ServerAddress string	`mapstructure:"SERVER_ADDRESS"`

}

func Load(path string) (config Config,error error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err !=nil {
		log.Print("Read Error",err)
		return
	}

	err = viper.Unmarshal(&config)
	
	return

}
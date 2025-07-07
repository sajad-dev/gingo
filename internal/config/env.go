package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var Config = &AppConfig{}

func BootConfig(address string) *AppConfig {
	err := godotenv.Load(address)
	if err != nil {
		panic(fmt.Sprintf("Config error : %s", err))
	}

	configSt := &AppConfig{}
	v := reflect.ValueOf(configSt).Elem()

	for i := 0; i < v.NumField(); i++ {
		env_value := os.Getenv(v.Type().Field(i).Name)
		field := v.Field(i)

		if env_value != "" {
			field.SetString(env_value)
		} else {
			field.SetString(v.Type().Field(i).Tag.Get("defualt"))
		}
	}

	Config = configSt

	return configSt
}

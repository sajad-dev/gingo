package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {	
	BootConfig("../../.env")	

	v := reflect.ValueOf(Config).Elem()

	for i := 0; i < v.NumField(); i++ {
		env_value := os.Getenv(v.Type().Field(i).Name)
		field := v.Field(i)
		if env_value != "" {
			assert.Equal(t,field.Interface() , env_value)
		} else {
			assert.Equal(t,v.Type().Field(i).Tag.Get("defuat") , env_value)
		}
	}

}

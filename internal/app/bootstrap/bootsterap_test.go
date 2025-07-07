package bootstrap

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootsterap (t *testing.T) {
	boot :=reflect.ValueOf(Counstructor(nil))

	for i:= 0 ; i < boot.NumField() ; i++{
		assert.NotNil(t,boot.Field(i).Interface())
	}
}

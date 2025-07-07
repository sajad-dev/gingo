package customerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomValidation(t *testing.T) {
	fe := FieldErrorCustom{
		ValidationTag:    "required",
		FieldNamespace:   "User.username",
		FieldValue:       "JohnDoe",
	}
	err := ConstructorError(fe)
	assert.Equal(t, err[0].Error(), "Error on field User.username: JohnDoe in validation tag required")
}

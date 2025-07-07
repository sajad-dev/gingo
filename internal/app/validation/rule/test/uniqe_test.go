package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/gingo/internal/app/validation/rule"
	"github.com/sajad-dev/gingo/internal/db/connection"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/sajad-dev/gingo/utils"
	"github.com/stretchr/testify/assert"
)

type User struct {
	ID    uint
	Email string 
}

func TestUniqeField(t *testing.T) {
	table.TablesVerfiy = []interface{}{&User{}}
	db := utils.SetupTestDB()
	connection.DB = db

	db.Create(&User{Email: "test@example.com"})

	validate := validator.New()
	validate.RegisterValidation("uniq", rule.UniqeField)

	type Input struct {
		Email string `validate:"uniq=user:email"`
	}

	input1 := Input{Email: "test@example.com"}
	err := validate.Struct(&input1)
	assert.Error(t, err)

	input2 := Input{Email: "new@example.com"}
	err = validate.Struct(&input2)
	assert.NoError(t, err)
}

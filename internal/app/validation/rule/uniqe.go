package rule

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/gingo/internal/db/connection"
	"github.com/sajad-dev/gingo/internal/db/table"
	"gorm.io/gorm"
)

func UniqeField (fl validator.FieldLevel) bool {
	tablesST := table.TablesVerfiy
	parms := strings.Split(fl.Param(),":") 
	if len(fl.Field().String()) == 0{
		return true
	}
	for _,value := range tablesST {
		re := reflect.TypeOf(value).Elem()
		if strings.ToLower(re.Name()) == parms[0]{
			err := connection.DB.Where(map[string]any{strings.ToLower(parms[1]):fl.Field().String()}).First(value)
			if errors.Is(err.Error,gorm.ErrRecordNotFound){
				return true
			}else{
				return false
			}
		}
	}
	return true
}

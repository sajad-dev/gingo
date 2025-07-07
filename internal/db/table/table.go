package table

import (
	"errors"

	"github.com/sajad-dev/gingo/internal/exception"
	"gorm.io/gorm"
)

var TablesVerfiy []any = []any{&Portfolio{},&Accounts{},&Projects{}}

func PublicMigration(db *gorm.DB) error {
	err := db.AutoMigrate(TablesVerfiy...)
	if err != nil {
		return errors.New(exception.Exception(err))
	}
	return nil
}

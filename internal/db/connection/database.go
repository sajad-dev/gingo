package connection

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/sajad-dev/gingo/internal/exception"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() error {
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", config.Config.DATABASE_USER, config.Config.DATABASE_PASSWORD, config.Config.DATABASE_ADDRESS, config.Config.DATABASE_PORT)

	sqlServer, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		return errors.New(exception.Exception(err))
	}

	_, err = sqlServer.Exec("CREATE DATABASE IF NOT EXISTS " + config.Config.DATABASE_NAME)
	if err != nil {
		return errors.New(exception.Exception(err))
	}

	dsnRoot = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.DATABASE_USER, config.Config.DATABASE_PASSWORD, config.Config.DATABASE_ADDRESS, config.Config.DATABASE_PORT, config.Config.DATABASE_NAME)
	db, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return errors.New(exception.Exception(err))
	}

	DB = db
	return nil
}

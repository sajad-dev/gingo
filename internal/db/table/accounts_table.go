package table

import (
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	FullName string `gorm:"not null;size:50"`
	Mobile   string
}



package bootstrap

import "gorm.io/gorm"

type Bootstrap struct {
	Auth      *AuthBoot
}

func Counstructor(db *gorm.DB) Bootstrap {
	return Bootstrap{
		Auth:      BootAuth(db),
	}
}

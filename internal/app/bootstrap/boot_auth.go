package bootstrap

import (
	"github.com/sajad-dev/gingo/internal/app/controller"
	"github.com/sajad-dev/gingo/internal/app/repository"
	"github.com/sajad-dev/gingo/internal/app/services"
	"gorm.io/gorm"
)

type AuthBoot struct {
	Controller controller.AuthHandller
}

func BootAuth(db *gorm.DB) *AuthBoot {
	repoAccounts := repository.AccountsRepo{DB: db}
	service := services.AuthService{RepoAccount: &repoAccounts}
	controller := controller.AuthController{Services: &service}

	return &AuthBoot{Controller: &controller}
}

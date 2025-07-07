package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo-helpers/utils"
	"github.com/sajad-dev/gingo/internal/app/repository"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/db/table"
)

type AuthService struct {
	RepoAccount repository.AccountsDataAccses
}

//go:generate mockgen -source=auth_service.go -destination=mocks/mock_srvices_auth.go -package=mocks
type AuthLogestic interface {
	Login(ctx *gin.Context, field validation.Login) (table.Accounts, string, error)
	Register(ctx *gin.Context, field validation.Register) (string, error)
}

var _ AuthLogestic = &AuthService{}

func (a *AuthService) Login(ctx *gin.Context, field validation.Login) (table.Accounts, string, error) {
	account, err := a.RepoAccount.GetUserByFields(ctx, map[string]any{"email": field.Email, "password": utils.PasswordHashing(field.Password)})
	if err != nil {
		return account, "", err
	}

	token, err := utils.CreateJWT(map[string]any{"email": account.Email, "id": account.ID})
	return account, token, err
}

func (a *AuthService) Register(ctx *gin.Context, field validation.Register) (string, error) {
	account := table.Accounts{Email: field.Email, FullName: field.FullName, Mobile: field.Mobile, Password: utils.PasswordHashing(field.Password)}
	err := a.RepoAccount.Create(ctx, account)
	if err != nil {
		return "", err
	}

	token, err := utils.CreateJWT(map[string]any{"email": account.Email, "id": account.ID})
	return token, err
}

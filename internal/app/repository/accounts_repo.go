package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/db/table"
	"gorm.io/gorm"
)

type AccountsRepo struct {
	DB *gorm.DB
}

//go:generate mockgen -source=accounts_repo.go -destination=mocks/mock_repo_accounts.go -package=mocks
type AccountsDataAccses interface {
	GetAll(*gin.Context) ([]table.Accounts, error)
	GetUserByFields(ctx *gin.Context, Qury map[string]any) (table.Accounts, error)
	Create(ctx *gin.Context, fields table.Accounts) error
	GetByID(ctx *gin.Context, id int) (table.Accounts, error)
	Update(ctx *gin.Context, field table.Accounts, id int) error
	Delete(ctx *gin.Context, id int) error
}

var _ AccountsDataAccses = &AccountsRepo{}

func (p *AccountsRepo) GetAll(ctx *gin.Context) ([]table.Accounts, error) {
	var accountsList []table.Accounts
	err := p.DB.Find(&accountsList).Error
	return accountsList, err
}

func (a *AccountsRepo) GetUserByFields(ctx *gin.Context, Qury map[string]any) (table.Accounts, error) {
	var user table.Accounts
	err := a.DB.Where(Qury).First(&user)
	return user, err.Error
}

func (a *AccountsRepo) Create(ctx *gin.Context, fields table.Accounts) error {
	return a.DB.Create(&fields).Error
}
func (p *AccountsRepo) GetByID(ctx *gin.Context, id int) (table.Accounts, error) {
	var AccountFirst table.Accounts
	err := p.DB.First(&AccountFirst, id).Error
	return AccountFirst, err
}
func (p *AccountsRepo) Update(ctx *gin.Context, field table.Accounts, id int) error {
	return p.DB.Model(&table.Accounts{}).Where("id = ?", id).Updates(field).Error
}

func (p *AccountsRepo) Delete(ctx *gin.Context, id int) error {
	return p.DB.Delete(&table.Accounts{}, id).Error
}


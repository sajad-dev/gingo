package test

// import (
// 	"testing"

// 	"github.com/sajad-dev/gingo/internal/app/repository"
// 	"github.com/sajad-dev/gingo/internal/db/table"
// 	"github.com/sajad-dev/gingo/utils"
// 	"github.com/stretchr/testify/assert"
// )

// var tableData = table.Portfolio{PortfolioName: "hihi"}

// func TestPortfolio_GetAll(t *testing.T) {
// 	db := utils.SetupTestDB()
// 	repo := &repository.PortfolioRepo{DB: db}

// 	repo.DB.Create(&tableData)
// 	data, err := repo.GetAll(nil)
// 	assert.NoError(t, err)
// 	assert.Equal(t, tableData.PortfolioName, data[0].PortfolioName)
// }

// func TestPortfolio_GetByID(t *testing.T) {
// 	db := utils.SetupTestDB()
// 	repo := &repository.PortfolioRepo{DB: db}

// 	repo.DB.Create(&table.Portfolio{})
// 	repo.DB.Create(&tableData)
// 	data, err := repo.GetByID(nil, 1)
// 	assert.NoError(t, err)
// 	assert.Equal(t, tableData.PortfolioName, data.PortfolioName)

// }

// func TestPortfolio_Create(t *testing.T) {
// 	db := utils.SetupTestDB()
// 	repo := &repository.PortfolioRepo{DB: db}

// 	err := repo.Create(nil, tableData)
// 	assert.NoError(t, err)

// 	var GetTestData table.Portfolio
// 	err = repo.DB.Last(&GetTestData).Error
// 	assert.NoError(t, err)
// 	assert.Equal(t, tableData.ID, GetTestData.ID)

// }

// func TestPortfolio_Update(t *testing.T) {
// 	db := utils.SetupTestDB()
// 	repo := &repository.PortfolioRepo{DB: db}

// 	repo.DB.Create(&tableData)

// 	err := repo.Update(nil, table.Portfolio{PortfolioName: "Haha is funny"}, int(tableData.ID))
// 	assert.NoError(t, err)

// 	var GetTestData table.Portfolio
// 	err = repo.DB.First(&GetTestData, tableData.ID).Error
// 	assert.NoError(t, err)
// 	data := tableData
// 	data.PortfolioName = "Haha is funny"
// 	assert.Equal(t, data.PortfolioName, GetTestData.PortfolioName)

// }

// func TestPortfolio_Delete(t *testing.T) {
// 	db := utils.SetupTestDB()
// 	repo := &repository.PortfolioRepo{DB: db}

// 	repo.DB.Create(&tableData)

// 	err := repo.Delete(nil, int(tableData.ID))
// 	assert.NoError(t, err)

// 	var GetTestData table.Portfolio
// 	err = repo.DB.Last(&GetTestData).Error
// 	assert.Error(t, err)

// }

// func TestPortfolio_GetLast(t *testing.T) {
// 	db := utils.SetupTestDB()
// 	repo := &repository.PortfolioRepo{DB: db}

// 	err := repo.Create(nil, tableData)
// 	assert.NoError(t, err)

// 	data,err := repo.GetLast(nil)
// 	assert.NoError(t, err)
// 	assert.Equal(t,tableData.ID,data.ID)


// }

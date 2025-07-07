package test

import (
	"testing"

	"github.com/sajad-dev/gingo-helpers/utils"
	"github.com/sajad-dev/gingo/internal/app/repository"
	"github.com/sajad-dev/gingo/internal/bootstrap"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	bootstrap.UtilsBoot()

	m.Run()
}

func TestAccountsRepo_GetAll(t *testing.T) {
	db := utils.SetupTestDB()
	repo := &repository.AccountsRepo{DB: db}

	account := table.Accounts{Email: "user@example.com", Password: "password123", FullName: "Test User"}
	db.Create(&account)

	result, err := repo.GetAll(nil)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "user@example.com", result[0].Email)
	assert.Equal(t, "Test User", result[0].FullName)
}

func TestAccountsRepo_GetByID(t *testing.T) {
	db := utils.SetupTestDB()
	repo := &repository.AccountsRepo{DB: db}

	// Insert test data
	account := table.Accounts{Email: "user@example.com", Password: "password123", FullName: "Test User"}
	db.Create(&account)

	// Fetch account by ID
	result, err := repo.GetByID(nil, int(account.ID))

	assert.NoError(t, err)
	assert.Equal(t, "user@example.com", result.Email)
	assert.Equal(t, "Test User", result.FullName)
}

func TestAccountsRepo_Create(t *testing.T) {
	db := utils.SetupTestDB()
	repo := &repository.AccountsRepo{DB: db}

	// Create a new account
	account := table.Accounts{Email: "newuser@example.com", Password: "password123", FullName: "New User"}
	err := repo.Create(nil, account)

	assert.NoError(t, err)

	// Verify the account was created
	var createdAccount table.Accounts
	db.First(&createdAccount, "email = ?", "newuser@example.com")
	assert.Equal(t, "newuser@example.com", createdAccount.Email)
	assert.Equal(t, "New User", createdAccount.FullName)
}

func TestAccountsRepo_Update(t *testing.T) {
	db := utils.SetupTestDB()
	repo := &repository.AccountsRepo{DB: db}

	// Insert test data
	account := table.Accounts{Email: "user@example.com", Password: "password123", FullName: "Test User"}
	db.Create(&account)

	// Update the account details
	account.FullName = "Updated User"
	err := repo.Update(nil, account, int(account.ID))

	assert.NoError(t, err)

	// Verify the account was updated
	var updatedAccount table.Accounts
	db.First(&updatedAccount, int(account.ID))
	assert.Equal(t, "Updated User", updatedAccount.FullName)
}

func TestAccountsRepo_Delete(t *testing.T) {
	db := utils.SetupTestDB()
	repo := &repository.AccountsRepo{DB: db}

	// Insert test data
	account := table.Accounts{Email: "user@example.com", Password: "password123", FullName: "Test User"}
	db.Create(&account)

	// Delete the account
	err := repo.Delete(nil, int(account.ID))

	assert.NoError(t, err)

	// Verify the account was deleted
	var deletedAccount table.Accounts
	result := db.First(&deletedAccount, int(account.ID))
	assert.Error(t, result.Error)
}
func TestAccountsRepo_GetUserByFields(t *testing.T) {
	db := utils.SetupTestDB()
	repo := &repository.AccountsRepo{DB: db}

	// Insert test data
	account := table.Accounts{
		Email:    "user@example.com",
		Password: "password123",
		FullName: "Test User",
		Mobile:   "123456789",
	}
	db.Create(&account)

	// Define the filter (query) to fetch user by email
	query := map[string]any{
		"email": "user@example.com",
	}

	// Fetch the user by fields (e.g., by email)
	result, err := repo.GetUserByFields(nil, query)

	assert.NoError(t, err)
	assert.Equal(t, "user@example.com", result.Email)
	assert.Equal(t, "Test User", result.FullName)
	assert.Equal(t, "123456789", result.Mobile)
}

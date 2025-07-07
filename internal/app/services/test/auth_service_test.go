package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sajad-dev/gingo/internal/app/repository/mocks"
	"github.com/sajad-dev/gingo/internal/app/services"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountsDataAccses(ctrl)

	loginField := validation.Login{
		Email:    "test@example.com",
		Password: "password123",
	}

	authService := &services.AuthService{
		RepoAccount: mockRepo,
	}

	expectedAccount := table.Accounts{
		Email: "test@example.com",
	}
	mockRepo.EXPECT().GetUserByFields(gomock.Any(), gomock.Any()).Return(expectedAccount, nil)


	account, _, err := authService.Login(nil, loginField)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)
}

func TestAuthService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountsDataAccses(ctrl)

	registerField := validation.Register{
		Email:    "newuser@example.com",
		FullName: "New User",
		Mobile:   "1234567890",
		Password: "newpassword123",
	}

	authService := &services.AuthService{
		RepoAccount: mockRepo,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)


	_, err := authService.Register(nil, registerField)

	assert.NoError(t, err)
}


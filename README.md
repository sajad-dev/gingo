# Gingo Framework

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.18+-00ADD8.svg?style=flat&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/license-MIT-green.svg?style=flat" alt="License">
  <img src="https://img.shields.io/badge/version-1.0.0-blue.svg?style=flat" alt="Version">
  <img src="https://img.shields.io/badge/status-stable-brightgreen.svg?style=flat" alt="Status">
</p>

<p align="center">
  <strong>A powerful and lightweight web framework with Go for building professional applications</strong>
</p>

---

## 📖 Table of Contents

- [Introduction](#-introduction)
- [Features](#-features)
- [Installation & Setup](#-installation--setup)
- [Quick Start](#-quick-start)
- [Project Structure](#-project-structure)
- [Complete Tutorial](#-complete-tutorial)
- [Practical Examples](#-practical-examples)
- [Testing](#-testing)
- [Documentation](#-documentation)

## 🚀 Introduction

**Gingo** is a modern and powerful web framework built on top of **Gin** and **GORM**, providing complete tools for building professional web applications. This framework is designed with clean architecture and modular structure to enable developers to easily create scalable projects.

### Gingo Ecosystem

| Package | Description | Installation |
|---------|-------------|--------------|
| **[gingo-cli](https://github.com/sajad-dev/gingo-cli)** | CLI tool for project management and scaffolding | `go install github.com/sajad-dev/gingo-cli@latest` |
| **[gingo-helpers](https://github.com/sajad-dev/gingo-helpers)** | Core library with helper functions | `go get github.com/sajad-dev/gingo-helpers@latest` |
| **[gear](https://github.com/sajad-dev/gear)** | Migration management and project execution tool | `go get github.com/sajad-dev/gear@latest` |

## ✨ Features

- 🔧 **Powerful CLI** - Complete command-line tools for project management
- 🗄️ **Migration System** - Database versioning management
- ✅ **Advanced Validation** - Strong validation with custom rules
- 🏗️ **Clean Architecture** - Clear separation of different layers
- 🔐 **Authentication** - Production-ready security system
- 📊 **Error Management** - Comprehensive error handling system
- 🧪 **Testing Support** - Built-in testing tools

## 📋 Prerequisites

- **Go 1.18+** - [Download Go](https://golang.org/dl/)
- **Database** - MySQL, PostgreSQL, or SQLite
- **Git** - For code management

## 🛠️ Installation & Setup

### Install Main CLI

```bash
# Install gingo-cli globally
go install github.com/sajad-dev/gingo-cli@latest

# Verify installation
gingo version
```

### Create New Project

```bash
# Method 1: Create new project
gingo new my-project
cd my-project

# Method 2: Use sample project
git clone https://github.com/sajad-dev/gingo-personal-backend-sample.git
cd gingo-personal-backend-sample
```

### Install Dependencies

```bash
# Install main packages
go get github.com/sajad-dev/gingo-helpers@latest
go get github.com/sajad-dev/gear@latest

# Install other dependencies
go mod tidy
```

## 🚦 Quick Start

### 1. Database Configuration

Edit the `internal/config/database.go` file:

```go
package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func DatabaseConfig() *gorm.Config {
    return &gorm.Config{
        // Database settings
    }
}

func DatabaseConnection() gorm.Dialector {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    return mysql.Open(dsn)
}
```

### 2. Run Migrations

```bash
# Create new migration
./gear make:migration create_users_table

# Run migrations
./gear migrate

# Rollback migration
./gear migrate:rollback
```

### 3. Start Server

```bash
# Run project
./gear run

# Or direct execution
go run cmd/main.go
```

## 🏗️ Project Structure

```
gingo/
├── 📁 build/                     # Build files
├── 📁 cmd/
│   └── main.go                   # Application entry point
├── 📁 gear/                      # CLI migration tool
├── 📁 integration/               # Integration tests
│   ├── auth_test.go
│   └── main_test.go
├── 📁 internal/
│   ├── 📁 app/                   # Application layer
│   │   ├── 📁 bootstrap/         # Application bootstrap
│   │   ├── 📁 controller/        # HTTP controllers
│   │   ├── 📁 repository/        # Data access layer
│   │   ├── 📁 response/          # Response models
│   │   ├── 📁 services/          # Business logic
│   │   ├── 📁 types/             # Data types
│   │   └── 📁 validation/        # Input validation
│   ├── 📁 bootstrap/             # Framework bootstrap
│   ├── 📁 config/                # Configuration management
│   ├── 📁 db/                    # Database connection
│   ├── 📁 exception/             # Error management
│   ├── 📁 middleware/            # HTTP middleware
│   ├── 📁 routes/                # Route definitions
│   └── 📁 server/                # HTTP server setup
├── 📁 storage/
│   └── 📁 files/                 # File storage
├── go.mod
├── go.sum
└── README.md
```

## 📚 Complete Tutorial

### Creating Controllers

```go
// internal/app/controller/user_controller.go
package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sajad-dev/gingo-helpers/response"
    "your-project/internal/app/services"
    "your-project/internal/app/validation"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{
        userService: userService,
    }
}

// List users
func (uc *UserController) Index(c *gin.Context) {
    users, err := uc.userService.GetAll()
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Error fetching users")
        return
    }
    
    response.Success(c, users, "Users retrieved successfully")
}

// Create new user
func (uc *UserController) Store(c *gin.Context) {
    var req validation.CreateUserRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        response.ValidationError(c, err)
        return
    }
    
    user, err := uc.userService.Create(req)
    if err != nil {
        response.Error(c, http.StatusBadRequest, err.Error())
        return
    }
    
    response.Success(c, user, "User created successfully")
}

// Show user
func (uc *UserController) Show(c *gin.Context) {
    id := c.Param("id")
    
    user, err := uc.userService.GetByID(id)
    if err != nil {
        response.Error(c, http.StatusNotFound, "User not found")
        return
    }
    
    response.Success(c, user, "User found successfully")
}

// Update user
func (uc *UserController) Update(c *gin.Context) {
    id := c.Param("id")
    var req validation.UpdateUserRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        response.ValidationError(c, err)
        return
    }
    
    user, err := uc.userService.Update(id, req)
    if err != nil {
        response.Error(c, http.StatusBadRequest, err.Error())
        return
    }
    
    response.Success(c, user, "User updated successfully")
}

// Delete user
func (uc *UserController) Delete(c *gin.Context) {
    id := c.Param("id")
    
    err := uc.userService.Delete(id)
    if err != nil {
        response.Error(c, http.StatusBadRequest, err.Error())
        return
    }
    
    response.Success(c, nil, "User deleted successfully")
}
```

### Creating Services

```go
// internal/app/services/user_service.go
package services

import (
    "errors"
    "your-project/internal/app/repository"
    "your-project/internal/app/types"
    "your-project/internal/app/validation"
)

type UserService struct {
    userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

func (us *UserService) GetAll() ([]types.User, error) {
    return us.userRepo.GetAll()
}

func (us *UserService) GetByID(id string) (*types.User, error) {
    if id == "" {
        return nil, errors.New("user ID is required")
    }
    
    return us.userRepo.GetByID(id)
}

func (us *UserService) Create(req validation.CreateUserRequest) (*types.User, error) {
    // Check for duplicate email
    existingUser, _ := us.userRepo.GetByEmail(req.Email)
    if existingUser != nil {
        return nil, errors.New("this email is already in use")
    }
    
    user := &types.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: req.Password, // Should be hashed in production
    }
    
    return us.userRepo.Create(user)
}

func (us *UserService) Update(id string, req validation.UpdateUserRequest) (*types.User, error) {
    user, err := us.GetByID(id)
    if err != nil {
        return nil, err
    }
    
    if req.Name != "" {
        user.Name = req.Name
    }
    if req.Email != "" {
        user.Email = req.Email
    }
    
    return us.userRepo.Update(user)
}

func (us *UserService) Delete(id string) error {
    user, err := us.GetByID(id)
    if err != nil {
        return err
    }
    
    return us.userRepo.Delete(user.ID)
}
```

### Creating Repositories

```go
// internal/app/repository/user_repository.go
package repository

import (
    "gorm.io/gorm"
    "your-project/internal/app/types"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (ur *UserRepository) GetAll() ([]types.User, error) {
    var users []types.User
    err := ur.db.Find(&users).Error
    return users, err
}

func (ur *UserRepository) GetByID(id string) (*types.User, error) {
    var user types.User
    err := ur.db.Where("id = ?", id).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*types.User, error) {
    var user types.User
    err := ur.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepository) Create(user *types.User) (*types.User, error) {
    err := ur.db.Create(user).Error
    return user, err
}

func (ur *UserRepository) Update(user *types.User) (*types.User, error) {
    err := ur.db.Save(user).Error
    return user, err
}

func (ur *UserRepository) Delete(id uint) error {
    return ur.db.Delete(&types.User{}, id).Error
}
```

### Defining Types

```go
// internal/app/types/user.go
package types

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    Name      string         `json:"name" gorm:"not null"`
    Email     string         `json:"email" gorm:"unique;not null"`
    Password  string         `json:"-" gorm:"not null"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
    return "users"
}
```

### Defining Validation

```go
// internal/app/validation/user_validation.go
package validation

type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=50" label:"Name"`
    Email    string `json:"email" validate:"required,email" label:"Email"`
    Password string `json:"password" validate:"required,min=8" label:"Password"`
}

type UpdateUserRequest struct {
    Name  string `json:"name" validate:"omitempty,min=2,max=50" label:"Name"`
    Email string `json:"email" validate:"omitempty,email" label:"Email"`
}
```

### Defining Routes

```go
// internal/routes/user_routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "your-project/internal/app/controller"
    "your-project/internal/middleware"
)

func UserRoutes(router *gin.Engine, userController *controller.UserController) {
    userGroup := router.Group("/api/users")
    {
        userGroup.GET("", userController.Index)
        userGroup.POST("", userController.Store)
        userGroup.GET("/:id", userController.Show)
        
        // Protected routes
        protected := userGroup.Use(middleware.AuthMiddleware())
        {
            protected.PUT("/:id", userController.Update)
            protected.DELETE("/:id", userController.Delete)
        }
    }
}
```

### Creating Middleware

```go
// internal/middleware/auth_middleware.go
package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/sajad-dev/gingo-helpers/response"
    "your-project/internal/app/services"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            response.Error(c, http.StatusUnauthorized, "Authorization token is required")
            c.Abort()
            return
        }
        
        token := strings.TrimPrefix(authHeader, "Bearer ")
        if token == authHeader {
            response.Error(c, http.StatusUnauthorized, "Invalid token format")
            c.Abort()
            return
        }
        
        // Validate token
        userID, err := services.ValidateToken(token)
        if err != nil {
            response.Error(c, http.StatusUnauthorized, "Invalid token")
            c.Abort()
            return
        }
        
        c.Set("user_id", userID)
        c.Next()
    }
}
```

## 🔧 Practical Examples

### Creating Migrations

```bash
# Create migration for users table
./gear make:migration create_users_table

# Create migration to add field
./gear make:migration add_phone_to_users_table

# Run migrations
./gear migrate

# Rollback latest migration
./gear migrate:rollback

# Reset all migrations
./gear migrate:reset
```

### Sample Migration File

```go
// migrations/2024_01_01_000000_create_users_table.go
package migrations

import (
    "gorm.io/gorm"
    "your-project/internal/app/types"
)

func up_create_users_table(db *gorm.DB) error {
    return db.AutoMigrate(&types.User{})
}

func down_create_users_table(db *gorm.DB) error {
    return db.Migrator().DropTable(&types.User{})
}
```

### Configuration Settings

```go
// internal/config/app.go
package config

import (
    "os"
    "strconv"
)

type AppConfig struct {
    Name        string
    Port        string
    Environment string
    Debug       bool
    JWTSecret   string
}

func LoadAppConfig() *AppConfig {
    debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
    
    return &AppConfig{
        Name:        getEnv("APP_NAME", "Gingo App"),
        Port:        getEnv("PORT", "8080"),
        Environment: getEnv("APP_ENV", "development"),
        Debug:       debug,
        JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
```

## 🧪 Testing

### Unit Tests

```go
// internal/app/services/user_service_test.go
package services

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "your-project/internal/app/validation"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetByEmail(email string) (*types.User, error) {
    args := m.Called(email)
    return args.Get(0).(*types.User), args.Error(1)
}

func TestUserService_Create(t *testing.T) {
    mockRepo := new(MockUserRepository)
    userService := NewUserService(mockRepo)
    
    req := validation.CreateUserRequest{
        Name:     "John Doe",
        Email:    "john@example.com",
        Password: "password123",
    }
    
    mockRepo.On("GetByEmail", req.Email).Return(nil, gorm.ErrRecordNotFound)
    mockRepo.On("Create", mock.AnythingOfType("*types.User")).Return(&types.User{}, nil)
    
    user, err := userService.Create(req)
    
    assert.NoError(t, err)
    assert.NotNil(t, user)
    mockRepo.AssertExpectations(t)
}
```

### Integration Tests

```go
// integration/user_test.go
package integration

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := setupRouter()
    
    user := map[string]interface{}{
        "name":     "John Doe",
        "email":    "john@example.com",
        "password": "password123",
    }
    
    jsonData, _ := json.Marshal(user)
    req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    
    assert.Equal(t, "User created successfully", response["message"])
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test ./integration/...

# Run specific tests
go test ./internal/app/services/...

# Show HTML coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🎯 Best Practices

### 1. Code Structure
- Use Clean Architecture
- Each layer should have its specific responsibility
- Use dependency injection

### 2. Error Management
- Always handle errors
- Error messages should be clear and understandable
- Use logging for error tracking

### 3. Security
- Never store passwords in plain text
- Use JWT for authentication
- Don't forget input validation

### 4. Performance
- Use connection pooling
- Optimize database queries
- Use caching when needed

## 📖 Documentation

### Useful Links

- **[Gin Documentation](https://gin-gonic.com/docs/)**
- **[GORM Documentation](https://gorm.io/docs/)**
- **[Go Documentation](https://golang.org/doc/)**
- **[Sample Project](https://github.com/sajad-dev/gingo-personal-backend-sample)**

### Help & Support

- **🐛 Bug Reports** - [Create Issue](https://github.com/sajad-dev/gingo/issues)
- **💡 Feature Requests** - [Request Feature](https://github.com/sajad-dev/gingo/issues)
- **❓ Questions** - [Discussions](https://github.com/sajad-dev/gingo/discussions)

## 🤝 Contributing

To contribute to this project:

1. **Fork** - Fork the repository
2. **Create Branch** - `git checkout -b feature/my-feature`
3. **Commit Changes** - `git commit -m 'Add some feature'`
4. **Push** - `git push origin feature/my-feature`
5. **Create Pull Request**

## 📄 License

This project is published under the **MIT** License - see the [LICENSE](LICENSE) file for more details.

---

<p align="center">
  Made with ❤️ by the Gingo Team
</p>

<p align="center">
  <a href="https://github.com/sajad-dev/gingo">⭐ Follow us on GitHub</a>
</p>

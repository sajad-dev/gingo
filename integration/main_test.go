package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/sajad-dev/gingo/internal/db/connection"
	"github.com/sajad-dev/gingo/internal/server"
	"github.com/sajad-dev/gingo/utils"
)

var engin *gin.Engine
var token string

func TestMain(m *testing.M) {
	// db := utils.SetupTestDB()
	db, resource, pool := SetupDB()
	defer pool.Purge(resource)

	connection.DB = db
	config.BootConfig("../.env")
	engin, _ = server.Http(8080, db)

	req := validation.Register{
		Email:           "info@sajad.com",
		Password:        "haha is funny",
		PasswordConfirm: "haha is funny",
	}

	res, _ := utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/register", Inputs: req, Engin: engin})

	var body map[string]string
	json.Unmarshal(res.Body.Bytes(), &body)
	token = body["token"]

	m.Run()
}

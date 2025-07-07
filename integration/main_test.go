package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo-helpers/core/bootstrap"
	"github.com/sajad-dev/gingo-helpers/types"
	"github.com/sajad-dev/gingo-helpers/utils"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/sajad-dev/gingo/internal/db/connection"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/sajad-dev/gingo/internal/server"
)

var engin *gin.Engine
var token string

func TestMain(m *testing.M) {
	bootstrap.Boot(types.Bootsterap{
		Config: types.ConfigUtils{
			STORAGE_PATH: config.Config.STORAGE_PATH,
			DATABASE:     table.TablesVerfiy,
			JWT:          config.Config.JWT,
		},
	})

	// db := utils.SetupTestDB()

	db, resource, pool := utils.SetupDB()
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

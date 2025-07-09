//go:build integration
// +build integration

package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sajad-dev/gingo-helpers/utils"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/stretchr/testify/assert"
)

func TestAuth_Register(t *testing.T) {
	req := validation.Register{}
	res, err := utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/register", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	assert.Equal(t, res.Result().StatusCode, 400)

	ok, err := utils.CheckValidationErr(res.Body.String(), "Email", "required")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = utils.CheckValidationErr(res.Body.String(), "Password", "required")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = utils.CheckValidationErr(res.Body.String(), "PasswordConfirm", "required")
	assert.NoError(t, err)
	assert.True(t, ok)

	req = validation.Register{
		Email:           "prj.sajad85@gmail.com",
		Password:        "haha is funny",
		PasswordConfirm: "haha is not funny :(",
	}
	res, err = utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/register", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	ok, err = utils.CheckValidationErr(res.Body.String(), "PasswordConfirm", "eqfield")
	assert.NoError(t, err)
	assert.True(t, ok)

	req = validation.Register{
		Email:           "prj.sajad85@gmail.com",
		Password:        "haha is funny",
		PasswordConfirm: "haha is funny",
	}
	res, err = utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/register", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	assert.Equal(t, 200, res.Result().StatusCode)
	t.Log(res.Body.String())

	var body map[string]string
	err = json.Unmarshal(res.Body.Bytes(), &body)
	assert.NoError(t, err)

	c, ok, err := utils.ValidJWT(body["token"])
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "prj.sajad85@gmail.com", (*c).Parameters["email"])

	req = validation.Register{
		Email:           "prj.sajad85@gmail.com",
		Password:        "haha is funny",
		PasswordConfirm: "haha is funny",
	}
	res, err = utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/register", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	assert.Equal(t, 400, res.Result().StatusCode)
	ok, err = utils.CheckValidationErr(res.Body.String(), "Email", "uniq")
	assert.NoError(t, err)
	assert.True(t, ok)

}

func TestAuth_Login(t *testing.T) {
	req := validation.Login{}
	res, err := utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/login", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	assert.Equal(t, res.Result().StatusCode, 400)

	ok, err := utils.CheckValidationErr(res.Body.String(), "Email", "required")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = utils.CheckValidationErr(res.Body.String(), "Password", "required")
	assert.NoError(t, err)
	assert.True(t, ok)

	req = validation.Login{
		Email:    "prj.sajadf85@gmail.com",
		Password: "haha is funny",
	}
	res, err = utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/login", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	req = validation.Login{
		Email:    "prj.sajad85@gmail.com",
		Password: "haha is funny",
	}
	res, err = utils.SendRequest(utils.Request{Method: http.MethodPost, Path: "/auth/login", Inputs: req, Engin: engin})
	assert.NoError(t, err)

	assert.Equal(t, 200, res.Result().StatusCode)

	var body map[string]any
	err = json.Unmarshal(res.Body.Bytes(), &body)
	assert.NoError(t, err)

	c, ok, err := utils.ValidJWT(body["token"].(string))
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "prj.sajad85@gmail.com", (*c).Parameters["email"])
}

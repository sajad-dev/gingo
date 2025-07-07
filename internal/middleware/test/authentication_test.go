package test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sajad-dev/gingo-helpers/utils"
// 	"github.com/sajad-dev/gingo/internal/middleware"
// 	"github.com/sajad-dev/gingo-helpers/utils"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAuth(t *testing.T){
	
// 	r := utils.CreateServer("/test","get",[]gin.HandlerFunc{middleware.AuthenticationMiddleware},8080)

// 	req, _ := http.NewRequest(http.MethodGet, "/test",nil)
// 	req.Header.Set("Authorization", "")
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t,w.Code,401)
	
// 	token ,_ := utils.CreateJWT(map[string]any{"message":"haha"})
// 	req, _ = http.NewRequest(http.MethodGet, "/test",nil)
// 	req.Header.Set("Authorization", token)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t,w.Code,200)
// }

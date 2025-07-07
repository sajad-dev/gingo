package exception

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestCallers(t *testing.T) {
	callers := Callers()

	assert.Equal(t, callers[0], "Called from testing.tRunner\n\t/usr/lib/go/src/testing/testing.go:1792\n")
	assert.Equal(t, callers[1], "Called from runtime.goexit\n\t/usr/lib/go/src/runtime/asm_amd64.s:1700\n")
}

func TestLogFile(t *testing.T) {

	DEBUG = "false"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DEBUG = "false"
		defer wg.Done()
		logFile()
		log.Println("Test print log ;)")
	}()
	wg.Wait()
	file, err := os.Open("../../storage/log/app.log")
	if err != nil {
		t.Fatalf("Error in opening file : %s", err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

	}

	stringLog := string(buf)
	assert.True(t, strings.Contains(stringLog, "Test print log ;)"))
}

func TestException(t *testing.T) {

	config.BootConfig("../../.env")
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	res := Exception(fmt.Errorf("This is a test error, don't be afraid. ;)"))

	w.Close()
	os.Stdout = origStdout
	var buf bytes.Buffer

	buf.ReadFrom(r)

	got := buf.String()
	assert.True(t, strings.Contains(res, got))
}
func TestServerException(t *testing.T) {
	config.BootConfig("../../.env")
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		err := fmt.Errorf("This is a test error, don't be afraid. ;)")
		ServerException(c,err)
	})

	DEBUG = "false"
	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"Server has error"`)

	DEBUG = "true"
	req, _ = http.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), "This is a test error, don't be afraid. ;)")

}



package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lerner17/shortener/internal/db"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSuccessRouter(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	database := db.GetInstance()
	var urlKeyValue map[string]string = make(map[string]string)
	urlList := []string{"https://youtube.com", "https://yandex.ru", "https://google.com", "https://go.dev"}

	for index := range urlList {
		dbID, _ := database.Insert(urlList[index])
		urlKeyValue[urlList[index]] = dbID
	}

	for key, value := range urlKeyValue {
		resp, _ := helpers.TestRequest(t, ts, "GET", fmt.Sprintf("/%s", value))
		assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
		assert.Equal(t, key, resp.Header.Get("Location"))
	}
}

func TestUndefinedRouter(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, _ := helpers.TestRequest(t, ts, "GET", fmt.Sprintf("/%s", "abc331"))
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

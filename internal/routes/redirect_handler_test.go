package routes

import (
	"fmt"
	"github.com/Lerner17/shortener/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp, string(respBody)
}

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
		resp, _ := testRequest(t, ts, "GET", fmt.Sprintf("/%s", value))
		assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
		assert.Equal(t, key, resp.Header.Get("Location"))
	}
}

func TestUndefinedRouter(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, _ := testRequest(t, ts, "GET", fmt.Sprintf("/%s", "abc331"))
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

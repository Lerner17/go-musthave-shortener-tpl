package handlers

import (
	"fmt"
	"github.com/Lerner17/shortener/cmd/shortener/db"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectHandler(t *testing.T) {
	database := db.GetInstance()
	var urlKeyValue map[string]string = make(map[string]string)
	urlList := []string{"https://youtube.com", "https://yandex.ru", "https://google.com", "https://go.dev"}

	type want struct {
		contentType string
		statusCode  int
	}

	for index := range urlList {
		dbID, _ := database.Insert(urlList[index])
		urlKeyValue[dbID] = urlList[index]
	}

	tests := []struct {
		name    string
		request string
		urlsMap map[string]string
		want    want
	}{
		{
			name:    "Success test #1",
			request: "/",
			urlsMap: urlKeyValue,
			want:    want{statusCode: 307, contentType: "plain/text"},
		},
		{
			name:    "Bad test #1",
			request: "/",
			urlsMap: map[string]string{"aaa": ""},
			want:    want{statusCode: 400, contentType: "plain/text"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, _ := range tt.urlsMap {
				var requestURL string = fmt.Sprintf("%s%s", tt.request, key)
				request := httptest.NewRequest(http.MethodGet, requestURL, nil)
				w := httptest.NewRecorder()
				h := http.HandlerFunc(RedirectHandler)
				h.ServeHTTP(w, request)
				result := w.Result()
				assert.Equal(t, tt.want.statusCode, result.StatusCode)
				assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			}

		})
	}
}

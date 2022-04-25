// Напишите сервис для сокращения длинных URL. Требования:
// Сервер должен быть доступен по адресу: http://localhost:8080.
// Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
// Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
// Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
// Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
var DataBase map[string]string = map[string]string{}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func generateShortUniqueID() string {
	for {
		var randomCandidate string = StringWithCharset(7, charset)
		if DataBase[randomCandidate] == "" {
			return randomCandidate
		}
	}
}

type CreateShortUrlBody struct {
	URL string `json:"url"`
}

func CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var body CreateShortUrlBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil || body.URL == "" {
		fmt.Println("ERROR!")
	}
	var generatedShortID string = generateShortUniqueID()
	DataBase[generatedShortID] = body.URL
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", generatedShortID)))
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	if fullURL := DataBase[url[1]]; len(url) > 1 && url[1] != "" && fullURL != "" {
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
	} else {
		w.Write([]byte("Undefined redirect"))
	}

}

func main() {

	http.HandleFunc("/createShortURL", CreateShortUrlHandler)
	http.HandleFunc("/", RedirectHandler)
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()

	log.Fatal(server.ListenAndServe())
}

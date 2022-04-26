package handlers

import (
	"encoding/json"
	"fmt"
	database "github.com/Lerner17/shortener/cmd/shortener/db"
	"net/http"
)

type createShortUrlBody struct {
	URL string `json:"url"`
}

func CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var body createShortUrlBody
	db := database.GetInstance()
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil || body.URL == "" {
		fmt.Println("ERROR!")
	}
	key, _ := db.Insert(body.URL)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", key)))
}

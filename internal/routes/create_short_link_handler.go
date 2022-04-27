package routes

import (
	"encoding/json"
	"fmt"
	database "github.com/Lerner17/shortener/internal/db"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	key, _ := db.Insert(body.URL)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", key)))
}

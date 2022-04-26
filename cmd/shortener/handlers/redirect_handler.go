package handlers

import (
	database "github.com/Lerner17/shortener/cmd/shortener/db"
	"net/http"
	"strings"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()
	url := strings.Split(r.URL.Path, "/")
	if fullURL, ok := db.Find(url[1]); len(url) > 1 && url[1] != "" && ok {
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
	} else {
		w.Write([]byte("Undefined redirect"))
	}

}

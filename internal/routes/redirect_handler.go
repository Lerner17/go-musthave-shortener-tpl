package routes

import (
	database "github.com/Lerner17/shortener/internal/db"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()
	urlID := chi.URLParam(r, "urlID")
	if fullURL, ok := db.Find(urlID); ok {
		w.Header().Set("Content-Type", "text/html")
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("URL id not found"))
	}

}

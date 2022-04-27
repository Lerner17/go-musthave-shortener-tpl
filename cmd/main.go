package main

import (
	"log"
	"net/http"

	"github.com/Lerner17/shortener/internal/routes"
)

func main() {
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

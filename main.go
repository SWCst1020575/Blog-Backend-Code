package main

import (
	"net/http"

	"Blog-Backend/cmd/api"

	"github.com/rs/cors"
)

func main() {
	router := api.NewRouter()
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8000", handler)

}

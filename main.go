package main

import (
	"net/http"
	"web_store/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.StartRoutes()
	http.ListenAndServe(":3000", nil)
}

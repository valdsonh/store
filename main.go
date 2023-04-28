package main

import (
	"net/http"
	"valdson/store/routes"
)

func main() {
	routes.LoadRouts()
	http.ListenAndServe(":8000", nil)
}

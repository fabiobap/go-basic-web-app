package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const HTTP_PORT = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port: %s", HTTP_PORT)
	_ = http.ListenAndServe(HTTP_PORT, nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const HTTP_PORT = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port: %s", HTTP_PORT)
	_ = http.ListenAndServe(HTTP_PORT, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error while parsing the template", err)
		return
	}

}

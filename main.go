package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/tacoroumen/Dynamic_website_golang/forms"
)

const addr = ":80" // TODO Make this configurable.

var templates *template.Template

func init() {
	templatesPath := filepath.Join("templates", "*.html")
	templates = template.Must(template.New("").ParseGlob(templatesPath))
}

func main() {

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/Aanvraag/", handleAanvraag)
	// Forms.go aanvraag gestart
	http.HandleFunc("/Aanvraag/reservering/", forms.HandleForms)
	http.HandleFunc("/Overons/", handleOverons)

	// Corrected ListenAndServe usage
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleAanvraag(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "aanvraag.html", nil)
	if err != nil {
		log.Println("error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleOverons(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "over_ons.html", nil)
	if err != nil {
		log.Println("error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

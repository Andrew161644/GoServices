package main

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func HelloPageHandler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "Go html",
		Message: "go htm is working",
	}
	tmpl, _ := template.ParseFiles("templates/main.html")
	tmpl.Execute(w, data)
}

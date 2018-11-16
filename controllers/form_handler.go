package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// FormHandler handles bird adding form
func FormHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tpl/form.gtpl")
	if err != nil {
		log.Fatal("Got error in template execute")
	}

	t.Execute(w, nil)
}

package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"text/template"
)

func main() {
	goji.Get("/", IndexHandler)
	goji.Serve()
}

func IndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

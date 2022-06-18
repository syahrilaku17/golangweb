package main

import (
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", handlerHello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	var address = "localhost:9000"
	http.ListenAndServe(address, nil)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}

}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"Title": "Hello, world!",
		"Body":  "This is my first Go web page.",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

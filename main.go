package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct{
	Item string
	Done bool
}

type PageData struct{
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request){
	data := PageData {
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Update news on X", Done: true},
			{Item: "Digest 3 posts on Medium", Done: true},
			{Item: "Learn Go for 1h", Done: false},
			{Item: "Update Solidity code", Done: false},
			{Item: "Go Swimming", Done: false},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("template/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	
	log.Fatal(http.ListenAndServe(":8080", mux))
}
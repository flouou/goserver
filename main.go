package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

//App loads the entire API set together for use.
func App() http.Handler {
	r := pat.New()

	r.Get("/", indexHandler)
	return r
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	p, _ := loadPage("index")
	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(res, p)
}

func main() {
	log.Print("Listening on Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", App()))
}

//Page contains the basic content of a Webpage
type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	templatename := title + ".html"
	body, _ := ioutil.ReadFile(templatename)
	return &Page{Title: title, Body: body}, nil
}

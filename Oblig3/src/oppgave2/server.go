package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"html/template"
	"fmt"
)


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/s/", pageServer)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//defines page
type Page struct {
	Title string
	Body []byte
}

//Saves body to a file with name Title
func (page *Page) save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)
}

//loads pages
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

//a handler for a scenario
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path[1:])
	page := "page" + request.URL.Path[1:2]
	http.ServeFile(writer, request, page)
}

func editHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/edit/"):]
	page := loadPage(title)
	renderTemplate(writer, "edit", page)
}

//handler for "/view/"
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	page := loadPage(title)
	renderTemplate(writer, "view", page)
}

//renders template
func renderTemplate(writer http.ResponseWriter, tmpl string, page *Page) {
	template, _ := template.ParseFiles(tmpl + ".html")
	template.Execute(writer, page)
}

//URL after '/s/' is the name of the html page returned
func pageServer (writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, request.URL.Path[len("/s/"):])
}


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/nei", nei)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/s/", pageServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//defines page
type Page struct {
	Title string
	Body []byte
}


//Saves body to a file with name Title
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//loads pages
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

//a handler for a scenario
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hellooo %s", r.URL.Path[1:])
	fmt.Fprintln(w, "Hello, client")
}

//handler for "nei"
func nei(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "NEI")
}

//handler for "/view/"
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):]
	//p := loadPage(title)
	w.Header().Set("lkol", "okk")
	apage,_ := ioutil.ReadFile("apage.txt")
	bpage := string(apage)
	fmt.Fprintf(w, bpage)
	//fmt.Fprintf(w, "<head><header>Heisann</header><link rel=\"spreadsheet\" type=\"text/css\" href=\"file:///C:/Users/Eier/Desktop/wiki.css\"></head><h1>ÆØÅ %s</h1><p>%s</p>", p.Title, p.Body)
}

//URL after '/s/' is the name of the html page returned
func pageServer (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}


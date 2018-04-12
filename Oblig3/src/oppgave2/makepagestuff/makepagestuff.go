package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
)

func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/s/", pageServer)
	http.HandleFunc("/edit/", editHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

//loads pages
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

func openJsonLocal() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Printf("Error: v%", err)
	}

	fmt.Printf("Successfully opened json\n")


	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}


	json.Unmarshal(byteValue, &users)
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
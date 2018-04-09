package main

import (
	"io/ioutil"
	"net/http"
	//"log"
	"html/template"
	"fmt"
	"os"
	"encoding/json"
)


func main() {
	jsontest()
	/*http.HandleFunc("/", handler)
	//http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/s/", pageServer)
	//http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))*/
}

//defines page
type Page struct {
	Title string
	Body []byte
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Username	string	`json:"username"`
	Address		Address	`json:"address"`
	Company		Company	`json:"company"`
	Phone		string	`json:"phone"`

}

type Address struct {
	Street	string	`json:"street"`
	Suite	string	`json:"suite"`

}

type Company struct {
	Name		string	`json:"name"`
	Catchphrase	string	`json:"catchPhrase"`
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

func jsontest() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Printf("Error: v%", err)
	}

	fmt.Printf("Successfully opened json: %s", jsonFile)


	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("\nUser ID: %d\n", users.Users[i].Id)
		fmt.Printf("User Name: %s\n", users.Users[i].Name)
		fmt.Printf("User Username: %s\n", users.Users[i].Username)
		fmt.Printf("Address: %s, %s\n", users.Users[i].Address.Street, users.Users[i].Address.Suite)
		fmt.Printf("Company: %s - %s\n", users.Users[i].Company.Name, users.Users[i].Company.Catchphrase)
		fmt.Printf("Phone: %s\n", users.Users[i].Phone)
	}
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


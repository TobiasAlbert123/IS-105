package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"html/template"
	"fmt"
	"encoding/json"
	"strconv"
	"time"
	"os"
)

func main() {
	go func() {
		http.HandleFunc("/", openJson)
	}()

	for i := 0; i < len(users.Users); i++ {
		jsonHandler(i)
	}
	fmt.Printf("\nMap contains %d entries, should be 60\n", len(jm))
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/view/", viewHandler)
	//http.HandleFunc("/s/", pageServer)
	//http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/1", makeJsonSite)
	http.HandleFunc("/a", attempt)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//defines page
type Page struct {
	Title string
	Body []byte
}

type jsonPerson struct {
	First	string
	Second	string
	Third	string
	Fourth	string
	Fifth	string
	Sixth	string
}

type Users struct {
	Users []User `json:" "`
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

//loads pages
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

/*
func loadPerson(index int) *jsonPerson {
	for i := 0; i < len(statNames); i++ {

	}
	first := statNames[1] + jm[statNames[1] + strconv.Itoa(index)]
	second := statNames[2] + jm[statNames[2] + strconv.Itoa(index)]
	third := statNames[3] + jm[statNames[3] + strconv.Itoa(index)]
	fourth := statNames[4] + jm[statNames[4] + strconv.Itoa(index)]
	fifth := statNames[5] + jm[statNames[5] + strconv.Itoa(index)]
	sixth := statNames[6] + jm[statNames[6] + strconv.Itoa(index)]

	return &jsonPerson{
		First: first,
		Second: second,
		Third: third,
		Fourth: fourth,
		Fifth: fifth,
		Sixth: sixth,
	}
}
*/
func jsontemp(writer http.ResponseWriter, person *jsonPerson) {
	temp, _ := template.ParseFiles("page-template.html")
	temp.Execute(writer, person)
}

func attempt(writer http.ResponseWriter, request *http.Request) {
		person := loadPerson(1)
		fmt.Println(person.Second)
		jsontemp(writer, person)
}

//a handler for a scenario
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello, client.")
}

var users Users

func openJsonLocal() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Printf("should not show Error: %v", err)
	}

	fmt.Printf("Successfully opened json\n")


	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}


	json.Unmarshal(byteValue, &users)
}
/*
func openJson(writer http.ResponseWriter, request *http.Request) {
	url := "https://jsonplaceholder.typicode.com/users"
	client := http.Client{
		Timeout: time.Second *2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("json url success")


	request.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(request)
	if getErr != nil {
		fmt.Println(err)
	}

	fmt.Println("json get success")

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(err)
	}

	fmt.Println("json read success")

	jsonErr := json.Unmarshal(body, &users)
	if jsonErr != nil {
		fmt.Println(err)
	}


	fmt.Println("json unmarshall success")

	fmt.Println(len(users.Users))

}
*/
var jm = make(map[string]string)

var statNames []string

func jsonHandler(index int) {


	//defines the name of the map elements
	element1 := "ID"
	element2 := "Name"
	element3 := "Username"
	element4 := "Address"
	element5 := "Company"
	element6 := "Phone"

	//prevents multiple values with same keys in map
	stringI := strconv.Itoa(index)

	//map elements created
	jm[element1 + stringI] = strconv.Itoa(users.Users[index].Id)
	jm[element2 + stringI] = users.Users[index].Name
	jm[element3 + stringI] = users.Users[index].Username
	jm[element4 + stringI] = users.Users[index].Address.Street + ", " + users.Users[index].Address.Suite
	jm[element5 + stringI] = users.Users[index].Company.Name + " - " + users.Users[index].Company.Catchphrase
	jm[element6 + stringI] = users.Users[index].Phone

	//adds element names to a slice for easier printing
	statNames = []string{element1, element2, element3, element4, element5, element6}

	fmt.Println()


	//prints a stat and the value (string) in console
	/*
	for i := 0; i < len(statNames); i++ {
		fmt.Println(statNames[i], jm[statNames[i]  + stringI])
	}*/

}

func makeJsonSite(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(users.Users); i++ {
		fmt.Fprintln(writer)
		for j := 0; j < len(statNames); j++ {
			fmt.Fprintf(writer, "%s: %s\n", statNames[j], jm[statNames[j] + strconv.Itoa(i)])
		}
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


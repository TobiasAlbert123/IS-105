package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"./pages/page1"
	"./pages/page2"
	"./pages/page3"
	"./pages/page4"
	"./pages/page5"
)

func main() {
	http.HandleFunc("/", helloClient)
	http.HandleFunc("/1", Page1)
	http.HandleFunc("/2", Page2)
	http.HandleFunc("/3", Page3)
	http.HandleFunc("/4", Page4)
	http.HandleFunc("/5", Page5)
	http.ListenAndServe(":8080", nil)
}

//for oppgave1
func helloClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, client.")
}


//defines the Template type, to be used in the html template
type Template struct {
	Title	string
	Name0	string
	Value0	string
	Name1	string
	Value1	string
	Name2	string
	Value2	string
	Name3	string
	Value3	string
	Name4	string
	Value4	string
	Name5	string
	Value5	string
	Name6	string
	Value6	string
}

//opens and returns the json data at url
func getJson(url string) []byte {
	client := http.Client{
		Timeout: time.Second *2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

//adds values to the Template type to make it ready for the html template
func loadTemplate(title string, names, values []string) *Template {
	//adds spacing between name and value for template
	//as well as filling in for missing or empty data
	for i := 0; i < len(names); i++ {
		names[i] += ": "
		if len(values[i]) == 0 {
			values[i] += "-missing-"
		}
	}

	//initialises an empty Template type, needs to happen within function or it will not clear properly
	t := Template{Title:title}
	t.Name0 = names[0]
	t.Value0 = values[0]
	//adds names and values to the Template type based on the names slice length
	for i := 0; i < len(names); i++ {
		switch i {
			case 1:
				t.Name1 = names[i]
				t.Value1 = values[i]
			case 2:
				t.Name2 = names[i]
				t.Value2 = values[i]
			case 3:
				t.Name3 = names[i]
				t.Value3 = values[i]
			case 4:
				t.Name4 = names[i]
				t.Value4 = values[i]
			case 5:
				t.Name5 = names[i]
				t.Value5 = values[i]
			case 6:
				t.Name6 = names[i]
				t.Value6 = values[i]
		}
	}
	return &t
}

//writes a header for the html page
//a separate template could be used, but probably too much effort for 3 lines
func writeTitle(w http.ResponseWriter, url string) {
	fmt.Fprintln(w, "<h1 style=\"font-size:3em; text-align:center;\">Datasets from: </h1>")
	fmt.Fprintf(w, "<p style=\"font-size:2em; text-align:center;\"><a href=\"%s\" target=\"_blank\">%s</a></p>\n",url, url)
	fmt.Fprint(w, "<br>")
}

//renders the template as a html page
func useTemplate(w http.ResponseWriter, page *Template) {
	template, _ := template.ParseFiles("page-template.html")
	template.Execute(w, page)
}

//all Page'x' functions are the same, but there must be some code duplication as
//they use a lot of different types which does not transfer well between functions
//PS: comments for pages are not duplicated
func Page1(w http.ResponseWriter, r *http.Request) {
	//initialises the page
	page  := page1.People{}
	//puts json data into the readable format defined in page
	jsonErr := json.Unmarshal(getJson(page1.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "a person"
	//writes header to the html page
	writeTitle(w, page1.Url)
	//uses template on each object in the json data
	for i := 0; i < len(page.People); i++ {
		names := []string{"Name", "Craft"}
		v := page.People[i]
		values := []string{v.Name, v.Craft}
		useTemplate(w, loadTemplate(title, names, values))
	}
}

func Page2(w http.ResponseWriter, r *http.Request) {
	page  := page2.Entries{}
	jsonErr := json.Unmarshal(getJson(page2.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "a county"
	writeTitle(w, page2.Url)
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"Kommune", "Fylke", "Navn"}
		v:= page.Entries[i]
		values := []string{v.Kommune, v.Fylke, v.Navn}
		useTemplate(w, loadTemplate(title, names, values))
	}
}

func Page3(w http.ResponseWriter, r *http.Request) {
	page  := page3.Entries{}
	jsonErr := json.Unmarshal(getJson(page3.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "a trip"
	writeTitle(w, page3.Url)
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"Kost", "Land", "Makssatser natt", "Verdensdel"}
		v := page.Entries[i]
		values := []string{v.Kost, v.Land, v.Makssatser_natt, v.Verdensdel}
		useTemplate(w, loadTemplate(title, names, values))
	}
}

func Page4(w http.ResponseWriter, r *http.Request) {
	page  := page4.Entries{}
	jsonErr := json.Unmarshal(getJson(page4.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "an organisation / thing"
	writeTitle(w, page4.Url)
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"Navn", "Organisasjonsform", "Stiftelsesdato", "Regdato", "Tlf", "Forretningsadresse", "Poststed"}
		v := page.Entries[i]
		values := []string{v.Navn, v.Organisasjonsform, v.Stiftelsesdato, v.Regdato, v.Tlf, v.Forretningsadresse, v.Poststed}
		useTemplate(w, loadTemplate(title, names, values))
	}
}

func Page5(w http.ResponseWriter, r *http.Request) {
	page := page5.Datasets{}
	jsonErr := json.Unmarshal(getJson(page5.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "a dataset"
	writeTitle(w, page5.Url)
	for i := 0; i < len(page.Datasets); i++ {
		names := []string{"ID", "Title", "Antall", "Beskrivelse"}
		v := page.Datasets[i]
		values := []string{v.Id, v.Title, v.Antall, v.Description[0].Value}
		useTemplate(w, loadTemplate(title, names, values))
	}
}
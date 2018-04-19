package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"./page1"
	"./page2"
	"./page3"
	"./page4"
	"./page5"
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

func helloClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, client.")
}

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

var p1 page1.People

var p2 page2.Entries

var p3 page3.Entries

var p4 page4.Entries

var p5 page5.Datasets


//opens and returns the json data at url
func getJson(url string) []byte {
	client := http.Client{
		Timeout: time.Second *2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}


func loadTemplate(index int, title string, names, values []string) *Template {
	for i := 0; i < len(names); i++ {
		names[i] += ": "
		if len(values[i]) == 0 {
			values[i] += "-missing-"
		}
	}

	t := Template{}
	t.Title = title
	for i := 0; i < index; i++ {
		switch i {
		case 0:
			t.Name0 = names[i]
			t.Value0 = values[i]
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

func useTemplate(writer http.ResponseWriter, page *Template) {
	template, _ := template.ParseFiles("page-template.html")
	template.Execute(writer, page)
}

func writeTitle(w http.ResponseWriter, url string) {
	fmt.Fprintln(w, "<h1 style=\"font-size:3em\">Datasets from: </h1>")
	fmt.Fprintf(w, "<a style=\"font-size:2em\" href=\"" + url + "\" target=\"_blank\">"+ url + "</a>")
	fmt.Fprint(w, "<br><br>")
}

func Page1(w http.ResponseWriter, r *http.Request) {
	page  := p1
	jsonErr := json.Unmarshal(getJson(page1.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "a person"
	writeTitle(w, page1.Url)
	for i := 0; i < len(page.People); i++ {
		names := []string{"Name", "Craft"}
		v := page.People[i]
		values := []string{v.Name, v.Craft}
		useTemplate(w, loadTemplate(len(names), title, names, values))
	}
}

func Page2(w http.ResponseWriter, r *http.Request) {
	page  := p2
	jsonErr := json.Unmarshal(getJson(page2.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "a set"
	writeTitle(w, page2.Url)
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"Kommune", "Fylke", "Navn"}
		v:= page.Entries[i]
		values := []string{v.Kommune, v.Fylke, v.Navn}
		useTemplate(w, loadTemplate(len(names), title, names, values))
	}
}

func Page3(w http.ResponseWriter, r *http.Request) {
	page  := p3
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
		useTemplate(w, loadTemplate(len(names), title, names, values))
	}
}

func Page4(w http.ResponseWriter, r *http.Request) {
	page  := p4
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
		useTemplate(w, loadTemplate(len(names), title, names, values))
	}
}

func Page5(w http.ResponseWriter, r *http.Request) {
	page := p5
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
		useTemplate(w, loadTemplate(len(names), title, names, values))
	}
}

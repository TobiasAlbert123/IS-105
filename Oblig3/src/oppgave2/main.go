package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"./page1"
	"./page5"
	"encoding/json"
	"html/template"
	"IS-105/Oblig3/src/oppgave2/page2"
	"IS-105/Oblig3/src/oppgave2/page3"
	"IS-105/Oblig3/src/oppgave2/page4"
)

func main() {
	http.HandleFunc("/1", Page1)
	http.HandleFunc("/2", Page2)
	http.HandleFunc("/3", Page3)
	http.HandleFunc("/4", Page4)
	http.HandleFunc("/5", Page5)
	http.ListenAndServe(":8080", nil)
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

var p2 page2.Title

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

func loadTemplate(title string, names, values []string) *Template {
	for i := 0; i < len(names); i++ {
		names[i] += ": "
		if len(values[i]) == 0 {
			values[i] += "-missing-"
		}
	}
	switch len(names) {
	case 2:
		return &Template{
			Title:  title,
			Name0:  names[0],
			Value0: values[0],
			Name1:  names[1],
			Value1: values[1],
		}
		break
	case 3:
		return &Template{
			Title:  title,
			Name0:  names[0],
			Value0: values[0],
			Name1:  names[1],
			Value1: values[1],
			Name2:  names[2],
			Value2: values[2],
		}
		break
	case 4:
		return &Template{
			Title:  title,
			Name0:  names[0],
			Value0: values[0],
			Name1:  names[1],
			Value1: values[1],
			Name2:  names[2],
			Value2: values[2],
			Name3: names[3],
			Value3: values[3],
		}
		break
	case 5:
		return &Template{
			Title:  title,
			Name0:  names[0],
			Value0: values[0],
			Name1:  names[1],
			Value1: values[1],
			Name2:  names[2],
			Value2: values[2],
			Name3: names[3],
			Value3: values[3],
			Name4: names[4],
			Value4: values[4],
		}
		break
	case 6:
		return &Template{
			Title:  title,
			Name0:  names[0],
			Value0: values[0],
			Name1:  names[1],
			Value1: values[1],
			Name2:  names[2],
			Value2: values[2],
			Name3: names[3],
			Value3: values[3],
			Name4: names[4],
			Value4: values[4],
			Name5: names[5],
			Value5: values[5],
		}
		break
	case 7:
		return &Template{
			Title:  title,
			Name0:  names[0],
			Value0: values[0],
			Name1:  names[1],
			Value1: values[1],
			Name2:  names[2],
			Value2: values[2],
			Name3: names[3],
			Value3: values[3],
			Name4: names[4],
			Value4: values[4],
			Name5: names[5],
			Value5: values[5],
			Name6: names[6],
			Value6: values[6],
		}
		break
	}
	return &Template{
		Title:title,
		Name0: names[0],
		Value0: values[0],
	}
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
		values := []string{page.People[i].Name, page.People[i].Craft}
		useTemplate(w, loadTemplate(title, names, values))
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
	for i := 0; i < 1; i++ {
		names := []string{"Label", ""}
		values := []string{page.Dataset.Dimension.Region.Label, ""}
		useTemplate(w, loadTemplate(title, names, values))
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
		useTemplate(w, loadTemplate(title, names, values))
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
		useTemplate(w, loadTemplate(title, names, values))
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
		useTemplate(w, loadTemplate(title, names, values))
	}
}

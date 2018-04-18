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
)

func main() {
	http.HandleFunc("/1", Page1)
	http.HandleFunc("/5", Page5)
	http.ListenAndServe(":8080", nil)
}

type Temp struct {
	Title	string
	Name1	string
	Value1	string
	Name2	string
	Value2	string
	Name3	string
	Value3	string
}

var p1 page1.People

var p2 string

var p3 string

var p4 string

var p5 page5.Datasets


//opens and returns the json data at url
func openJson(url string) []byte {
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

func loadTemp(index int, names, values []string) *Temp {
	name1 := names[0]
	value1 := values[0]
	name2 := names[1]
	value2 := values[1]
	if len(names) > 2 {
		name3 := names[2]
		value3 := values[2]

		return &Temp {
			Name1:name1,
			Value1:value1,
			Name2:name2,
			Value2:value2,
			Value3: value3,
			Name3: name3,
		}
	}
	return &Temp {
		Name1:name1,
		Value1:value1,
		Name2:name2,
		Value2:value2,
	}
}

func useTemplate(writer http.ResponseWriter, page *Temp) {
	template, _ := template.ParseFiles("page-template.html")
	template.Execute(writer, page)
}

func Page1(w http.ResponseWriter, r *http.Request) {
	page  := p1
	jsonErr := json.Unmarshal(openJson(page1.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	for i := 0; i < len(page.People); i++ {
		names := []string{"Name", "Craft"}
		values := []string{page.People[i].Name, page.People[i].Craft}
		useTemplate(w, loadTemp(i, names, values))
	}
}

func Page5(w http.ResponseWriter, r *http.Request) {
	page := p5
	jsonErr := json.Unmarshal(openJson(page5.Url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	for i := 0; i < len(page.Datasets); i++ {
		names := []string{"ID", "Title", "Antall: "}
		values := []string{page.Datasets[i].Id, page.Datasets[i].Title, page.Datasets[i].Antall}
		useTemplate(w, loadTemp(i, names, values))
	}
}

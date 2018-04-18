package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"./tes"
	"./page5"
	"html/template"
)

func main() {
	http.HandleFunc("/1", Page1)
	http.HandleFunc("/2", Page2)
	http.HandleFunc("/3", Page3)
	http.HandleFunc("/4", Page4)
	http.HandleFunc("/5", Page5)
	http.ListenAndServe(":8080", nil)
}

type Temp struct {
	Name1 string
	Value1 string
	Name2 string
	Value2 string
	Name3 string
	Value3 string
}

var page1 tes.People

var page3 = tes.Tes

var page55 page5.Datasets


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

func Page1(w http.ResponseWriter, r *http.Request) {
	page := page1
	//unmarshal json and puts it into page
	jsonErr := json.Unmarshal(openJson(tes.Uuu), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	for i := 0; i < len(page.People); i++ {
		a := page.People[i]
		fmt.Fprintf(w,"Name: %s\n", a.Name)
		fmt.Fprintf(w,"Craft: %s\n", a.Craft)
	}
	tes.Bbb()
}

func Page3(w http.ResponseWriter, r *http.Request) {
	page := page3
	//unmarshal json and puts it into page
	jsonErr := json.Unmarshal(openJson(tes.Uuu), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	for i := 0; i < len(page.People); i++ {
		a := page.People[i]
		fmt.Fprintf(w,"Name: %s\n", a.Name)
		fmt.Fprintf(w,"Craft: %s\n", a.Craft)
	}
}

func Page2(w http.ResponseWriter, r *http.Request) {
	page := page55
	//unmarshal json and puts it into page
	url := "https://data.norge.no/api/dcat/data.json?page=1"
	jsonErr := json.Unmarshal(openJson(url), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	for i := 0; i < len(page.Datasets); i++ {
		a := page.Datasets[i]
		fmt.Fprintf(w,"\nID: %s\n", a.Id)
		fmt.Fprintf(w,"Title: %s\n", a.Title)
		fmt.Fprintf(w,"Antall: %s\n", a.Antall)
		fmt.Fprintf(w,"Distribution: %s\n", a.Distribution)
		fmt.Fprintf(w,"Description: %s\n", a.Description)
	}
}

func loadPerson(index int) *tes.Person{
	person := page1.People[index]
	name := person.Name
	craft := person.Craft
	return &tes.Person{
		Name: name,
		Craft: craft,
	}
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

func Page4(w http.ResponseWriter, r *http.Request) {
	url := "https://data.norge.no/api/dcat/data.json?page=1"
	jsonErr := json.Unmarshal(openJson(url), &page55)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	for i := 0; i < len(page55.Datasets); i++ {
		names := []string{"ID", "Title", "Antall: "}
		values := []string{page55.Datasets[i].Id, page55.Datasets[i].Title, page55.Datasets[i].Antall}
		useTemplate(w, loadTemp(i, names, values))
	}
}

func Page5(w http.ResponseWriter, r *http.Request) {
	jsonErr := json.Unmarshal(openJson(tes.Uuu), &page1)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	for i := 0; i < len(page1.People); i++ {
		//useTemplate(w, loadPerson(i))
		names := []string{"Name", "Craft"}
		values := []string{page1.People[i].Name, page1.People[i].Craft}
		useTemplate(w, loadTemp(i, names, values))
	}

}

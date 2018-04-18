package page5

import (
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	openJson()
	http.HandleFunc("/5", printToServer)
	http.ListenAndServe(":8080", nil)
}

type Datasets struct {
	Datasets		[]Dataset `json:"datasets"`
}

type Dataset struct {
	Id			string	`json:"id"`
	Title		string	`json:"title"`
	Antall		string		`json:"antall"`
	Description	[]Description	`json:"description"`
	Distribution []Distribution	`json:"distribution"`
}

type Distribution struct {
	Id		string	`json:"id"`
	Title	string	`json:"title"`
}

type Description struct {
	Language	string	`json:"language"`
	Value	string	`json:"value"`
}

var datasets Datasets

func openJson() {
	url := "https://data.norge.no/api/dcat/data.json?page=1"

	client := http.Client{
		Timeout: time.Second *2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {

	}

	response, err := client.Do(request)
	if err != nil {

	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}

	jsonerr := json.Unmarshal(body, &datasets)
	if jsonerr != nil {

	}


	fmt.Printf("Length: %d\n", len(datasets.Datasets))

	for i := 0; i < len(datasets.Datasets); i++ {
		printAll(i)

	}
}

func printAll(index int) {
	d:= datasets.Datasets[index]
	fmt.Printf("\nID: %s\n", d.Id)
	fmt.Printf("Tittel: %s\n", d.Title)
	fmt.Printf("Antall: %s\n", d.Antall)
	if len(d.Description) > 0 {
		s := d.Description[0].Value
		s = strings.Replace(s, "\u003Cp\u003E", "", -1)
		s = strings.Replace(s, "\u003C/p\u003E", "", -1)
		fmt.Printf("Beskrivelse: %s\n", s)
		fmt.Println()
	}
}

func printToServer(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(datasets.Datasets); i++ {
		d := datasets.Datasets[i]
		fmt.Fprintf(writer,"\nID: %s\n", d.Id)
		fmt.Fprintf(writer,"Tittel: %s\n", d.Title)
		fmt.Fprintf(writer,"Antall: %s\n", d.Antall)
		if len(d.Description) > 0 {
			fmt.Fprintf(writer, "Beskrivelse: %s\n", d.Description[0].Value)
		}
	}
}



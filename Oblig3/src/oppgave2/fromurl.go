package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Peoples struct {
	Peoples	[]People	`json:"people"`
}


type People struct {
	Name	string	`json:"name"`
	Craft	string	`json:"craft"`
}

var peoples Peoples

func main() {

	url := "http://api.open-notify.org/astros.json" //"https://jsonplaceholder.typicode.com/users"

	aClient := http.Client {
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}


	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := aClient.Do(req)
	if getErr != nil {
		fmt.Println(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(err)
	}

	peoples = Peoples{}
	jsonErr := json.Unmarshal(body, &peoples)
	if jsonErr != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(peoples.Peoples); i++ {
		fmt.Printf("Name - %s\nCraft - %s\n\n", peoples.Peoples[i].Name, peoples.Peoples[i].Craft)
	}

	http.HandleFunc("/4", four)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func four (writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(peoples.Peoples); i++ {
		fmt.Fprintf(writer, "Name - %s\nCraft - %s\n\n", peoples.Peoples[i].Name, peoples.Peoples[i].Craft)
	}
}

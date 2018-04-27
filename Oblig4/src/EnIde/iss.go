package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"runtime"
	"os/exec"
)

type Location struct {
	IssPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

func main () {
	location := Location{}

	//mapsUrl := "https://www.google.com/maps/search/18.6785+89.7448"
	for {
		err := json.Unmarshal(getJson(), &location)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Unix timestamp: %d\n", location.Timestamp)
		fmt.Printf("'Normal' timestamp: %s\n", time.Unix(int64(location.Timestamp), 0))
		fmt.Printf("Latitude: %s\n", location.IssPosition.Latitude)
		fmt.Printf("Longitude: %s\n\n", location.IssPosition.Longitude)
		time.Sleep(time.Second *5)
		lat := location.IssPosition.Latitude
		long := location.IssPosition.Longitude
		mapsUrl := "https://www.google.com/maps/search/" + lat + "+" + long
		fmt.Println(mapsUrl)
		open(mapsUrl)
		//placeholder := "https://www.google.com/maps/search/-50.1085+-154.3239"
	}
	//mapsBaseUrl := "https://www.google.com/maps/search/" + "lat" + "+" + "long"
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

// embeddedlinkbase = "https://www.google.com/maps/embed/v1/place?key=AIzaSyAx9uiZK2gNb3oNORe0-SLxO72f8-NYlaI&q=" + lat + "+" + long

func getJson() []byte {
	url := "http://api.open-notify.org/iss-now.json"
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

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
	"html/template"
)

type Location struct {
	IssPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Country	string
	UpdateFrequency	int
}

type CountryFinder struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
	}
	Status	string	`json:"status"`
	Country	string
}

//not needed as var, but need to store it somewhere
var apiKey = "AIzaSyAx9uiZK2gNb3oNORe0-SLxO72f8-NYlaI"

var geoKey = "AIzaSyDh4iNsKY2S8cT-qrwjkDZENR2fgo4oDvY"

var geoKey2 = "AIzaSyCXfJONlkP8c1PM0ZBOJFBtFR7hRhapQQw"

var sliceOfGeoKeys = []string{geoKey, geoKey2}

var geoKeysUsed = 0

var currentGeoKey = sliceOfGeoKeys[geoKeysUsed]

var updateFrequency = 10

var invalidData = false

func main () {
	//overloadCountryFinder("AIzaSyCXfJONlkP8c1PM0ZBOJFBtFR7hRhapQQw")
	http.HandleFunc("/", Page)
	http.HandleFunc("/err", ErrPage)
	http.HandleFunc("/2", Page2)
	http.ListenAndServe(":8080", nil)
	//location := Location{}

	//mapsUrl := "https://www.google.com/maps/search/18.6785+89.7448"
	/*
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
		//open(mapsUrl)

		//placeholder := "https://www.google.com/maps/search/-50.1085+-154.3239"
	}*/
	//mapsBaseUrl := "https://www.google.com/maps/search/" + "lat" + "+" + "long"
}

func getCountry(lat, long string) string{
	results := CountryFinder{}
	//url := "https://maps.googleapis.com/maps/api/geocode/json?latlng=-48.489638,22.767979&key=AIzaSyDh4iNsKY2S8cT-qrwjkDZENR2fgo4oDvY"
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng="+lat+","+long+"&key="+geoKey
	err := json.Unmarshal(getJson(url), &results)
	if err != nil {
		log.Fatal(err)
	}
	switch results.Status {
		case "OVER_QUERY_LIMIT":
			nextGeoKey()
			getCountry(lat, long)
		case "ZERO_RESULTS":
			results.Country = "Location is not in a country (it is probably in the ocean)"
			fmt.Println(results.Country)
		default:
			country := results.Results[0].AddressComponents[len(results.Results[0].AddressComponents)-1].LongName
			results.Country = country
			fmt.Println("country: ", country)
	}
	return results.Country

}

//moves on to next key for finding country, if available
func nextGeoKey() {
	geoKeysUsed++
	if geoKeysUsed <= len(sliceOfGeoKeys) {
		currentGeoKey = sliceOfGeoKeys[geoKeysUsed]
	} else {
		log.Fatal("No more API keys for finding country available")
	}
}

//used to find error messages when api key usage has been exceeded
func overloadCountryFinder(overloadKey string) {
	for i := 0; i < 2500; i++ {
		ting, _ := http.Get("https://maps.googleapis.com/maps/api/geocode/json?latlng=-18.489638,22.767979&key="+overloadKey)
		read, _ := ioutil.ReadAll(ting.Body)
		fmt.Println(string(read))
		fmt.Println(i)
	}
}

//unmarshals json
func unmarshalJson() *Location{
	location := Location{}
	url := "http://api.open-notify.org/iss-now.json"
	err := json.Unmarshal(getJson(url), &location)
	if err != nil {
		log.Fatal(err)
	}
	//printJson(location)
	if isEmpty(location) {
		invalidData = true
	}
	location.Country = getCountry(location.IssPosition.Latitude, location.IssPosition.Longitude)
	location.UpdateFrequency = updateFrequency
	return &location
}

func isEmpty(location Location) bool {
	//location.Timestamp = 0
	if location.Timestamp == 0  || location.IssPosition.Latitude == "" || location.IssPosition.Longitude == "" {
		return true
	}
	return false
}

//-28.489638,22.767979

//prints json in console
func printJson(location Location) {
	fmt.Printf("Unix timestamp: %d\n", location.Timestamp)
	fmt.Printf("'Normal' timestamp: %s\n", time.Unix(int64(location.Timestamp), 0))
	fmt.Printf("Latitude: %s\n", location.IssPosition.Latitude)
	fmt.Printf("Longitude: %s\n\n", location.IssPosition.Longitude)
	time.Sleep(time.Second *5)
	lat := location.IssPosition.Latitude
	long := location.IssPosition.Longitude
	mapsUrl := "https://www.google.com/maps/search/" + lat + "+" + long
	fmt.Println(mapsUrl)
}

//renders template to client
func renderTemplate(w http.ResponseWriter, page *Location) {
	t, err := template.ParseFiles("iss.html")
	if err != nil {
		log.Fatal(err)
	}
	if !invalidData {
		t.Execute(w, page)
	} else {
		fmt.Fprintln(w, "404: Some or all data were nil")
	}
}

//opens a url in default browser
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

//gets json from url
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

//main page
func Page(w http.ResponseWriter, r *http.Request) {
	//for {
		renderTemplate(w, unmarshalJson())
		//checkForErrors(w, "http://127.0.0.1:8080/")
		time.Sleep(time.Second * time.Duration(updateFrequency))
		//http.ServeFile(w, r, "empty.html")
		//fmt.Fprintln(w, "<script>document.getElementById('body').innerHTML = '';</script>")
	//}
}

//page for testing errors
func ErrPage(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "iss2.html")
		time.Sleep(time.Second * 1)
		checkForErrors(w, "http://127.0.0.1:8080/2")
}

//another random page for errors
func Page2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "iss2.html")
}

//checks errors
func checkForErrors(w http.ResponseWriter, url string) {

	/*
		page, err := http.Get(url)
		if err != nil {
			fmt.Println("Error getting page at url")
			log.Fatal(err)
		}
		reading, err := ioutil.ReadAll(page.Body)
		if err != nil {
			fmt.Println("Error reading page body")
			log.Fatal(err)
		}
		fmt.Fprintln(w, string(string(reading)))
	if strings.Contains(string(reading), "AIzaSyAx9uiZK2gNb3oNORe0-SLxO72f8-NYlaI") {
		fmt.Fprintln(w,"=======================================================")
		fmt.Fprintln(w,"You dont have the correct API key")
		fmt.Fprintln(w,"=======================================================")
	}
	if strings.Contains(string(reading), "Lat") {
		fmt.Fprintln(w,"=======================================================")
		fmt.Fprintln(w,"=======================================================")
	}*/

}
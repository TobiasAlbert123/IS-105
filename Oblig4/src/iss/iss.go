package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"html/template"
	"math"
	"runtime"
	"os/exec"
	"strconv"
)

func main() {
	invalidData = false
	http.HandleFunc("/", Page)
	http.HandleFunc("/css", CssPage)
	http.HandleFunc("/s", StaticPage)
	http.ListenAndServe(":8080", nil)
}

func runForTest() {
	go main()
	time.Sleep(time.Second*10)
	open("http://localhost:8080/")
	time.Sleep(time.Second*5)
}

/*
*STRUCTS
 */

//struct for issData Position API
type issData struct {
	Pos struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
	Message         string `json:"message"`
	UnixTime        int    `json:"timestamp"`
	UserTime        time.Time
	LocalTime		time.Time
	LocalTimeMsg	string
	Country         string
	UpdateFrequency int
	Name0           string
	Name1           string
	Name2           string
	Name3           string
	Name4           string
	Name5           string
	DayA            int
	DayB            int
	HourA           int
	HourB           int
	TimeZone		string
	Elevation		int
	ElevationMessage	string
}

//struct for the Google Reverse Geocoding API
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
	ErrorMessage	string	`json:"error_message"`
}

type TimeZoneFinder struct {
	DstOffset    int    `json:"dstOffset"`
	RawOffset    int    `json:"rawOffset"`
	Status       string `json:"status"`
	TimeZoneID   string `json:"timeZoneId"`
	TimeZoneName string `json:"timeZoneName"`
	ErrorMessage	string	`json:"error_message"`
}

type ElevationFinder struct {
	Results []struct {
		Elevation float64 `json:"elevation"`
		Location  struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Resolution float64 `json:"resolution"`
	} `json:"results"`
	Status string `json:"status"`
	ErrorMessage	string	`json:"error_message"`
}

/*
VARIABLES
 */

//not needed as a global var, but need to store it somewhere
//API Key for Google Maps Embed, limited at something like 2 million requests per day
var apiKey = "AIzaSyAx9uiZK2gNb3oNORe0-SLxO72f8-NYlaI"

//these are API Keys for reverse geocoding, timezones and elevation
//limited at 2500 requests per day per key
var geoKey = "AIzaSyDh4iNsKY2S8cT-qrwjkDZENR2fgo4oDvY"
var geoKey2 = "AIzaSyCXfJONlkP8c1PM0ZBOJFBtFR7hRhapQQw"
var geoKey3 = "AIzaSyC-eVqkAILczjX3KpepxwR2cAyXgaEvb8c"
var geoKey4 = "AIzaSyBSsIw_sA9wGy-ptDdUZ8pxI5elu9azujs"
var SliceOfGeoKeys = []string{geoKey, geoKey2, geoKey3, geoKey4}

var GeoKeysUsed = 0

var currentGeoKey = SliceOfGeoKeys[GeoKeysUsed]

//amount of seconds
var updateFrequency = 15

//changed to true when the iss api is invalid
var invalidData bool

//var outside recursive function to prevent it from resetting infinitely
var attemptedUnmarshals = 0

//this is used during testing
var globalError = ""

/*
JSON HANDLING
 */

//gets json from url
func getJson(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		globalError += "Error at http.Get\n"
		//log.Fatal("Error at http.Get")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		globalError += "Error at ioutil.ReadAll\n"
		//log.Fatal("Error at ioutil.Readall")
	}
	return body
}

//unmarshals and formats json
func formatJson() *issData {
	iss := issData{}
	url := "http://api.open-notify.org/iss-now.json"
	err := json.Unmarshal(getJson(url), &iss)
	if err != nil {
		globalError += "Error at json unmarshal\n"
		//log.Fatal(err)
	}
	//printJson(iss)
	if isEmpty(iss) {
		invalidData = true
	}
	//country information
	iss.Country = getCountry(iss.Pos.Latitude, iss.Pos.Longitude)

	//how often the page updates
	iss.UpdateFrequency = updateFrequency

	//time astronauts have been in space
	names, days, hours := astronautInfo()
	iss.Name0, iss.Name1, iss.Name2, iss.Name3, iss.Name4, iss.Name5 = names[0], names[1], names[2], names[3], names[4], names[5]
	iss.DayA, iss.DayB = days[0], days[1]
	iss.HourA, iss.HourB = hours[0], hours[1]

	//timestamp at user location
	iss.UserTime = time.Unix(int64(iss.UnixTime), 0)

	//timezone and local time
	timezone, id := getTimeZone(iss.Pos.Latitude, iss.Pos.Longitude, iss.UnixTime)
	iss.TimeZone = timezone
	if id != "N/A" {
		zone, _ := time.LoadLocation(id)
		iss.LocalTime = iss.UserTime.In(zone)
	}

	//checks that there is a timezone to print to page
	if timezone != "N/A" {
		iss.LocalTimeMsg = "Timestamp (local timezone): " + iss.LocalTime.String()
	} else {
		iss.LocalTimeMsg = ""
	}

	//elevation and ocean depth
	elevation := int(getElevation(iss.Pos.Latitude, iss.Pos.Longitude))

	if elevation < 0 {
		iss.ElevationMessage = "Ocean depth: "
	} else {
		iss.ElevationMessage = "Elevation: "
	}
	iss.ElevationMessage += strconv.Itoa(elevation) + "m"

	//attempts the function again 10 times if data collection was unsuccessful
	if iss.Message != "success" && attemptedUnmarshals < 10 {
		fmt.Println(iss.Message, len(iss.Message))
		//global var to avoid it resetting when function restarts
		attemptedUnmarshals++
		if attemptedUnmarshals < 10 {
			time.Sleep(time.Millisecond * 5)
			formatJson()
		}
		//log.Fatalf("%d attemps at collecting ISS API data were unsuccessful. Ensure url (%s) is correct", attemptedUnmarshals, url)
	} else if iss.Message == "" { //needed when testing API link or program breaks for some reason
		return &iss
	}
	return &iss
}

/*
GOOGLE APIS
 */

//returns the country at lat, long
func getCountry(lat, long string) string{
	results := CountryFinder{}
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng="+lat+","+long+"&key="+currentGeoKey
	//testurl := "https://maps.googleapis.com/maps/api/geocode/json?latlng=51.6204,-60.6336&key=AIzaSyCXfJONlkP8c1PM0ZBOJFBtFR7hRhapQQw"

	err := json.Unmarshal(getJson(url), &results)
	if err != nil {
		//log.Fatal(err)
	}
	switch results.Status {
		case "OVER_QUERY_LIMIT":
			nextGeoKey()
			getCountry(lat, long)
			results.Country = "N/A - API Keys depleted"
		case "ZERO_RESULTS":
			results.Country = "N/A"
		case "INVALID_REQUEST":
			results.Country = "N/A"
			globalError += "Invalid request for geocoding\n"
		case "REQUEST_DENIED":
			results.Country = "N/A"
			globalError += "request denied for country: " + results.ErrorMessage +"\n"
		case "OK":
			//country := results.Results[0].AddressComponents[len(results.Results[0].AddressComponents)-1].LongName
			country := ""
			for i := 0; i < len(results.Results[0].AddressComponents); i++ {
				for j := 0; j < len(results.Results[0].AddressComponents[i].Types); j++ {
					if results.Results[0].AddressComponents[i].Types[j] == "country" {
						country = results.Results[0].AddressComponents[i].LongName
					}
				}
			}
			results.Country = country
		default:
			fmt.Println("Country status: ", results.Status)
	}
	return results.Country

}

//returns timezone and offsett from unixtime (in seconds)
func getTimeZone(lat, long string, unixTime int) (string, string){
	timestamp := strconv.Itoa(unixTime)
	//testurl := "https://maps.googleapis.com/maps/api/timezone/json?location=41.634663,-111.189675&timestamp=1525377238&key=AIzaSyDh4iNsKY2S8cT-qrwjkDZENR2fgo4oDvY"
	url := "https://maps.googleapis.com/maps/api/results/json?location="+lat+","+long+"&timestamp="+timestamp+"&key="+currentGeoKey
	results := TimeZoneFinder{}
	err := json.Unmarshal(getJson(url), &results)
	if err != nil {
		//log.Fatal("Unmarshal error: ", err)
	}

	switch results.Status {
	case "OVER_QUERY_LIMIT":
		nextGeoKey()
		getTimeZone(lat, long, unixTime)
		results.TimeZoneName = "N/A - API Keys depleted"
		results.TimeZoneID = "N/A - API Keys depleted"
	case "ZERO_RESULTS":
		results.TimeZoneName = "N/A"
		results.TimeZoneID = "N/A"
	case "INVALID_REQUEST":
		results.TimeZoneName = "N/A"
		results.TimeZoneID = "N/A"
		globalError += "Invalid request for timezone\n"
	case "REQUEST_DENIED":
		results.TimeZoneName = "N/A"
		results.TimeZoneID = "N/A"
		globalError += "request denied for timezone\n"
		fmt.Println(results.ErrorMessage)
	case "OK":
		break
	default:
		fmt.Println("Timezone status: ", results.Status, results.TimeZoneName)
	}
	return results.TimeZoneName, results.TimeZoneID
}

//returns elevation or ocean depth
func getElevation(lat, long string) float64 {
	url := "https://maps.googleapis.com/maps/api/results/json?locations="+lat+","+long+"&key="+currentGeoKey
	//testurl := "https://maps.googleapis.com/maps/api/elevation/json?locations=-51.3502,64.5989&key=AIzaSyDgGyEYCnYDCWCtODdiM-DSnuUTcN2XKCo"

	results := ElevationFinder{}
	err := json.Unmarshal(getJson(url), &results)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}

	//fmt.Println(results.Results[0].Elevation)
	switch results.Status {
	case "OVER_QUERY_LIMIT":
		nextGeoKey()
		getElevation(lat, long)
		results.Results[0].Elevation = 0
	case "ZERO_RESULTS":
		results.Results[0].Elevation = 0
	case "INVALID_REQUEST":
		globalError += "Invalid request for elevation\n"
	case "REQUEST_DENIED":
		globalError += "request denied for elevation\n"
		fmt.Println(results.ErrorMessage)
	case "OK":
		break
	default:
		fmt.Println("Elevation status", results.Status)
	}
	return results.Results[0].Elevation
}

/*
OTHER FUNCTIONS
 */

//moves on to next key for finding country, if available
func nextGeoKey() {
	GeoKeysUsed++
	if GeoKeysUsed <= len(SliceOfGeoKeys) {
		currentGeoKey = SliceOfGeoKeys[GeoKeysUsed]
	} else {
		//log.Fatal("No more API keys for reverse geocoding available")
	}
}

//checks if iss api is invalid
func isEmpty(iss issData) bool {
	//iss.UnixTime = 0
	if iss.UnixTime == 0  || iss.Pos.Latitude == "" || iss.Pos.Longitude == "" {
		return true
	}
	return false
}

//renders template to client
func renderTemplate(w http.ResponseWriter, page *issData) {
	t, err := template.ParseFiles("iss.html")
	if err != nil {

	}
	if !invalidData {
		t.Execute(w, page)

	} else {
		fmt.Fprintln(w, "404: Some or all data were nil. Sorry!")
	}
}

//returns the names of the astronauts and how long they've been in space
func astronautInfo() ([]string, []int, []int){
	//names of the astronauts at ISS
	names := []string{"Scott Tingle", "Anton Shkaplerov", "Norishige Kanai", "Andrew Feustel", "Richard R. Arnold", "Oleg Artemyev"}

	//Sets types Time at entered dates, which corresponds to the launch dates
	//source for dates: https://www.worldspaceflight.com/bios/currentlyinspace.php
	launchA := time.Date(2017, time.December, 17, 7, 21, 0, 0, time.UTC)
	launchB := time.Date(2018, time.March, 21, 17, 44, 0, 0, time.UTC)

	//time since launches in hours
	timeSinceA := time.Since(launchA).Hours()
	timeSinceB := time.Since(launchB).Hours()

	//hours since launch (type float64) divided by 24 to get days, then floors the number and convert to type int
	daysA := int(math.Floor(timeSinceA/24))
	daysB := int(math.Floor(timeSinceB/24))

	//same as above, but '% 24' takes the remainder after dividing by 24
	hoursA := int(math.Floor(timeSinceA)) % 24
	hoursB := int(math.Floor(timeSinceB)) % 24

	//slices to be returned
	days := []int{daysA, daysB}
	hours := []int{hoursA, hoursB}
	
	return names, days, hours
}

/*
PAGES
 */

//main page
func Page(w http.ResponseWriter, r *http.Request) {
	//renders the page
	renderTemplate(w, formatJson())
	//prints to console which API key is being used, #4 is the last one
	fmt.Printf("Using API Key #%d\n", GeoKeysUsed+1)
	//sleeps for updatefrequency
	time.Sleep(time.Second * time.Duration(updateFrequency))
	http.ServeFile(w, r, "empty.html")
	Page(w, r)
}

func StaticPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, formatJson())
}

//page where css file is hosted (used by template in RunServer page)
func CssPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "css.css")
}

//opens a url in default browser
//use for test purposes
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

/**
*UNUSED FUNCTIONS
 */

//used to find error messages when api key usage has been exceeded
func overloadCountryFinder(overloadKey string) {
	for i := 0; i < 2500; i++ {
		ting, _ := http.Get("https://maps.googleapis.com/maps/api/geocode/json?latlng=-18.489638,22.767979&key="+overloadKey)
		read, _ := ioutil.ReadAll(ting.Body)
		fmt.Println(string(read))
		fmt.Println(i)
	}
}
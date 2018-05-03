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
}

type TimeZoneFinder struct {
	DstOffset    int    `json:"dstOffset"`
	RawOffset    int    `json:"rawOffset"`
	Status       string `json:"status"`
	TimeZoneID   string `json:"timeZoneId"`
	TimeZoneName string `json:"timeZoneName"`
}

//not needed as var, but need to store it somewhere
//API Key for Google Maps Embed
var apiKey = "AIzaSyAx9uiZK2gNb3oNORe0-SLxO72f8-NYlaI"

//these are API Keys for reverse geocoding
var geoKey = "AIzaSyDh4iNsKY2S8cT-qrwjkDZENR2fgo4oDvY"
var geoKey2 = "AIzaSyCXfJONlkP8c1PM0ZBOJFBtFR7hRhapQQw"
var geoKey3 = ""
var geoKey4 = ""
var geoKey5 = ""
var sliceOfGeoKeys = []string{geoKey, geoKey2}

var geoKeysUsed = 0

var currentGeoKey = sliceOfGeoKeys[geoKeysUsed]

var updateFrequency = 10

var invalidData bool

func main () {
	//overloadCountryFinder("AIzaSyCXfJONlkP8c1PM0ZBOJFBtFR7hRhapQQw")
	invalidData = false
	http.HandleFunc("/", Page)
	http.HandleFunc("/css", CssPage)
	http.ListenAndServe(":8080", nil)
	time.Sleep(time.Minute*10)
	//location := issData{}

	//mapsUrl := "https://www.google.com/maps/search/18.6785+89.7448"
	/*
	for {
		err := json.Unmarshal(getJson(), &location)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Unix timestamp: %d\n", location.UnixTime)
		fmt.Printf("'Normal' timestamp: %s\n", time.Unix(int64(location.UnixTime), 0))
		fmt.Printf("Latitude: %s\n", location.Pos.Latitude)
		fmt.Printf("Longitude: %s\n\n", location.Pos.Longitude)
		time.Sleep(time.Second *5)
		lat := location.Pos.Latitude
		long := location.Pos.Longitude
		mapsUrl := "https://www.google.com/maps/search/" + lat + "+" + long
		fmt.Println(mapsUrl)
		//open(mapsUrl)

		//placeholder := "https://www.google.com/maps/search/-50.1085+-154.3239"
	}*/
	//mapsBaseUrl := "https://www.google.com/maps/search/" + "lat" + "+" + "long"
}

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


var attemptedUnmarshals = 0

//unmarshals and formats json
func formatJson() *issData {
	iss := issData{}
	url := "http://api.open-notify.org/iss-now.json"
	err := json.Unmarshal(getJson(url), &iss)
	if err != nil {
		log.Fatal(err)
	}
	//printJson(iss)
	if isEmpty(iss) {
		invalidData = true
	}

	iss.Country = getCountry(iss.Pos.Latitude, iss.Pos.Longitude)
	iss.UpdateFrequency = updateFrequency
	names, days, hours := timeSinceLaunch()
	iss.Name0, iss.Name1, iss.Name2, iss.Name3, iss.Name4, iss.Name5 = names[0], names[1], names[2], names[3], names[4], names[5]
	iss.DayA, iss.DayB = days[0], days[1]
	iss.HourA, iss.HourB = hours[0], hours[1]
	iss.UserTime = time.Unix(int64(iss.UnixTime), 0)

	//timezone here

	//attempts the function again 10 times if data collection was unsuccessful
	if iss.Message != "success" && attemptedUnmarshals < 10 {
		fmt.Println(iss.Message,len(iss.Message))
		attemptedUnmarshals++
		if attemptedUnmarshals < 10 {
			time.Sleep(time.Millisecond*5)
			formatJson()
		}
		log.Fatalf("%d attemps at collecting ISS API data were unsuccessful. Ensure url (%s) is correct", attemptedUnmarshals, url)
	}
	return &iss
}

func getCountry(lat, long string) string{
	results := CountryFinder{}
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng="+lat+","+long+"&key="+currentGeoKey

	//url for a location in Canada that lists json data slightly differently
	//url := "https://maps.googleapis.com/maps/api/geocode/json?latlng=51.6204,-60.6336&key=AIzaSyDh4iNsKY2S8cT-qrwjkDZENR2fgo4oDvY"

	err := json.Unmarshal(getJson(url), &results)
	if err != nil {
		log.Fatal(err)
	}
	switch results.Status {
		case "OVER_QUERY_LIMIT":
			nextGeoKey()
			getCountry(lat, long)
		case "ZERO_RESULTS":
			results.Country = "Country information unavailable - location likely in the ocean"
		default:
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
	}
	return results.Country

}

func getTimeZone(lat, long string, unixTime int) (string, int){
	timestamp := strconv.Itoa(unixTime)
	url := "https://maps.googleapis.com/maps/api/timezone/json?location="+lat+","+long+"&timestamp="+timestamp+"&key="+currentGeoKey
	timezone := TimeZoneFinder{}
	err := json.Unmarshal(getJson(url), &timezone)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}


	switch timezone.Status {
	case "OVER_QUERY_LIMIT":
		nextGeoKey()
		getTimeZone(lat, long, unixTime)
	case "ZERO_RESULTS":
		timezone.TimeZoneName = "Timezone not found"
	}
	return timezone.TimeZoneName, timezone.RawOffset
}

//moves on to next key for finding country, if available
func nextGeoKey() {
	geoKeysUsed++
	if geoKeysUsed <= len(sliceOfGeoKeys) {
		currentGeoKey = sliceOfGeoKeys[geoKeysUsed]
	} else {
		log.Fatal("No more API keys for reverse geocoding available")
	}
}

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
		log.Fatal(err)
	}
	if !invalidData {
		t.Execute(w, page)
	} else {
		fmt.Fprintln(w, "404: Some or all data were nil. Sorry!")
	}
}

func timeSinceLaunch() ([]string, []int, []int){
	names := []string{"Scott Tingle", "Anton Skhaplerov", "Norishige Kanai", "Andrew Feustel", "Richard Arnold", "Oleg Martemyev"}

	//Sets types Time at entered dates, which corresponds to the launch dates
	//source for dates: https://www.worldspaceflight.com/bios/currentlyinspace.php
	launchA := time.Date(2017, time.December, 17, 7, 21, 0, 0, time.UTC)
	launchB := time.Date(2018, time.March, 21, 17, 44, 0, 0, time.UTC)

	//time since launches in hours
	timeSinceA := time.Now().Sub(launchA).Hours()
	timeSinceB := time.Now().Sub(launchB).Hours()


	//hours since launch (type float64) divided by 24 to get days, then floors the number and convert to type int
	daysA := int(math.Floor(timeSinceA/24))
	daysB := int(math.Floor(timeSinceB/24))

	//same as above, but '% 24' takes the remainder after dividing by 24
	hoursA := int(math.Floor(timeSinceA)) % 24
	hoursB := int(math.Floor(timeSinceB)) % 24

	//slices to be returned
	days := []int{daysA, daysB}
	hours := []int{hoursA, hoursB}

	ticker0 := makeTicker()
	ticker1 := makeTicker()
	ticker2 := makeTicker()
	ticker3 := makeTicker()
	ticker4 := makeTicker()
	ticker5 := makeTicker()

	tickers := []<-chan time.Time{ticker0}
	//false to prevent it from running (instead of putting code in a comment)
	if true {
			for {
				select {
				case <-tickers[0]:
					t := time.Now()
					elapsed := t.Sub(launchA)
					fmt.Printf("%s has been in space for %.0f days \n", names[0], elapsed.Hours()/24)
				/*
				case <-ticker0:
					t := time.Now()
					elapsed := t.Sub(launchA)
					fmt.Printf("%s has been in space %.0f days\n", names[0], elapsed.Hours()/24)*/

				case <-ticker1:
					t := time.Now()
					elapsed := t.Sub(launchA)
					fmt.Printf("%s has been in space %.0f days\n", names[1], elapsed.Hours()/24)

				case <-ticker2:
					t := time.Now()
					elapsed := t.Sub(launchA)
					fmt.Printf("%s has been in space %.0f days\n", names[2], elapsed.Hours()/24)

				case <-ticker3:
					t := time.Now()
					elapsed := t.Sub(launchB)
					fmt.Printf("%s has been in space %.0f days\n", names[3], elapsed.Hours()/24)

				case <-ticker4:
					t := time.Now()
					elapsed := t.Sub(launchB)
					fmt.Printf("%s has been in space %.0f days\n", names[4], elapsed.Hours()/24)

				case <-ticker5:
					t := time.Now()
					elapsed := t.Sub(launchB)
					fmt.Printf("%s has been in space %.0f days\n", names[5], elapsed.Hours()/24)

				}
			}
	}
	return names, days, hours
}

func makeTicker() <-chan time.Time{
	time.Sleep(time.Millisecond*50)
	ticker := time.Tick(time.Second*5)
	return ticker
}

//main page
func Page(w http.ResponseWriter, r *http.Request) {
	for {
		renderTemplate(w, formatJson())
		//checkForErrors(w, "http://127.0.0.1:8080/")
		time.Sleep(time.Second * time.Duration(updateFrequency))
		//http.ServeFile(w, r, "empty.html")
		/* Alternatively:
		fmt.Fprintln(w, "<script>document.getElementById('body').innerHTML = '';</script>")
		*/
		//time.Sleep(time.Millisecond * 1000)
	}
}

//page where css file is hosted (used by template in main page)
func CssPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "css.css")
}

/**
*Functions not in use below
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

//prints json in console
func printJson(iss issData) {
	fmt.Printf("Unix timestamp: %d\n", iss.UnixTime)
	fmt.Printf("'Normal' timestamp: %s\n", time.Unix(int64(iss.UnixTime), 0))
	fmt.Printf("Latitude: %s\n", iss.Pos.Latitude)
	fmt.Printf("Longitude: %s\n\n", iss.Pos.Longitude)
	time.Sleep(time.Second *5)
	lat := iss.Pos.Latitude
	long := iss.Pos.Longitude
	mapsUrl := "https://www.google.com/maps/search/" + lat + "+" + long
	fmt.Println(mapsUrl)
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

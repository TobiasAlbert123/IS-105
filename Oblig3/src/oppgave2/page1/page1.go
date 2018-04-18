package page1

type People struct {
	People	[]Person	`json:"people"`
}

type Person struct {
	Name	string	`json:"name"`
	Craft	string	`json:"craft"`
}

var Url = "http://api.open-notify.org/astros.json"
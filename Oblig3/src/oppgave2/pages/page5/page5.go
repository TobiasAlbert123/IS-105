package page5

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

type Description struct {
	Language	string	`json:"language"`
	Value	string	`json:"value"`
}

type Distribution struct {
	Id		string	`json:"id"`
	Title	string	`json:"title"`
}

var Url = "https://data.norge.no/api/dcat/data.json?page=1"



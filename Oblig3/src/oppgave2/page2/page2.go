package page2

type Entries struct {
	Entries	[]Entry	`json:"entries"`
}

type Entry struct {
	Kommune	string	`json:"kommune"`
	Fylke	string	`json:"fylke"`
	Navn	string	`json:"navn"`
}


var Url = "https://hotell.difi.no/api/json/difi/geo/kommune"
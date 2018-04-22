package page4

type Entries struct {
	Entries	[]Entry	`json:"entries"`
}

type Entry struct {
	Stiftelsesdato		string	`json:"stiftelsesdato"`
	Organisasjonsform	string	`json:"organisasjonsform"`
	Navn				string	`json:"navn"`
	Tlf					string	`json:"tlf"`
	Forretningsadresse	string	`json:"forretningsadr"`
	Poststed			string	`json:"forradrpoststed"`
	Regdato				string	`json:"regdato"`
}

var Url = "https://hotell.difi.no/api/json/brreg/enhetsregisteret?page=8"
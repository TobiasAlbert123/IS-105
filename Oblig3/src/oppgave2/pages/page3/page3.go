package page3

type Entries struct {
	Entries	[]Entry	`json:"entries"`
}

type Entry struct {
	Kost	string	`json:"kost"`
	Land	string	`json:"land"`
	Makssatser_natt	string	`json:"makssatser natt"`
	Verdensdel	string	`json:"verdensdel"`
}

var Url = "https://hotell.difi.no/api/json/fad/reise/utland?"
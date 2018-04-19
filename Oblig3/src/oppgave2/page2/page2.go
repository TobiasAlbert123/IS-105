package page2

type Title struct {
	Dataset	Dataset	`json:"dataset"`
}

type Dataset struct {
	Dimension	Dimension	`json:"dimension"`
}

type Dimension struct {
	Region	Region	`json:"Region"`
}

type Region struct {
	Label	string	`json:"label"`
	Category	Category	`json:"category"`
}

type Category struct {

}



var Url = "https://data.ssb.no/api/v0/dataset/1104.json?lang=no"
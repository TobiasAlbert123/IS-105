package tes

import (
	"fmt"
)

func Bbb() {
	fmt.Println()
}

type People struct {
	People	[]Person	`json:"people"`
}


type Person struct {
	Name	string	`json:"name"`
	Craft	string	`json:"craft"`
}

var Tes People


var Uuu = "http://api.open-notify.org/astros.json"
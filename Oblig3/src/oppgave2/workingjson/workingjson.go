package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)


func main() {
	jsontest()
}
type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Username	string	`json:"username"`
	Address		Address	`json:"address"`
	Company		Company	`json:"company"`
	Phone		string	`json:"phone"`
}

type Address struct {
	Street	string	`json:"street"`
	Suite	string	`json:"suite"`

}

type Company struct {
	Name		string	`json:"name"`
	Catchphrase	string	`json:"catchPhrase"`
}

func jsontest() {
	//opens json
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Printf("Error: v%", err)
	}

	fmt.Printf("Successfully opened json: %s", jsonFile)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("\nUser ID: %d\n", users.Users[i].Id)
		fmt.Printf("User Name: %s\n", users.Users[i].Name)
		fmt.Printf("User Username: %s\n", users.Users[i].Username)
		fmt.Printf("Address: %s, %s\n", users.Users[i].Address.Street, users.Users[i].Address.Suite)
		fmt.Printf("Company: %s - %s\n", users.Users[i].Company.Name, users.Users[i].Company.Catchphrase)
		fmt.Printf("Phone: %s\n", users.Users[i].Phone)
	}
}


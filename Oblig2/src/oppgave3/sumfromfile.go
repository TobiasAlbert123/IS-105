/*package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Oppgave 3")
	file, err := os.Open("Oblig2/src/oppgave3/testfile.txt")
	if err != nil {

	}
	buf := make([]byte, 1024)
	fmt.Println(file.Read(buf))
}*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

func addup(tall, tall2 int) int{
	return tall + tall2
}

func main() {
	fmt.Println("Oppgave 3")
	file, err := ioutil.ReadFile("Oblig2/src/oppgave3/testfile.txt")
	if err != nil {

	}
	fileString := string(file)
	splitString := []string(strings.Split(fileString, ""))
	var ints [2]int
	for i := 0; i < len(splitString); i++ {
		ints[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			fmt.Println("int conversion error")
		}
	}
	fmt.Printf("Total is %d\n", addup(ints[0], ints[1]))
	//ints, err := ReadInts(strings.NewReader(file))
	//fmt.Println(ints, err)
	//was looking at https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
}

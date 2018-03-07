package notmain

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
)

func addup(tall, tall2 int) int{
	return tall + tall2
}

func WriteSum(number int) {
	file, err := os.Create(testfile)
	if err != nil {
		fmt.Println("Error creating file")
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d", number))
	if err != nil {
		fmt.Println("Error writing")
	} else {
		fmt.Printf("Wrote sum %d to %s\n", number, testfile)
	}
}

func SumFromFile() int{
	fmt.Println("Oppgave 3")
	file, err := ioutil.ReadFile(testfile)
	if err != nil {

	}
	fileString := string(file)
	splitString := []string(strings.Split(fileString, "\n"))
	var ints [2]int
	for i := 0; i < len(splitString); i++ {
		ints[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			fmt.Println("int conversion error")
		}
	}
	total := addup(ints[0], ints[1])
	fmt.Printf("Total is %d\n", total)
	return total
}

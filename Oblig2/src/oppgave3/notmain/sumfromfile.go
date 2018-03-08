package notmain

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
)

//sums numbers from file
func SumFromFile() int{
	file, err := ioutil.ReadFile(testfile)
	if err != nil {

	}
	fileString := string(file)
	splitString := []string(strings.Split(fileString, "\n"))
	var ints [2]int
	for i := 0; i < len(splitString); i++ {
		ints[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			fmt.Println("Could not convert numbers in file to int")
			os.Exit(1)
		}
	}
	total := ints[0] + ints[1]
	return total
}

//writes sum to the file
func WriteSum(number int) {
	file, err := os.Create(testfile)
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d", number))
	if err != nil {
		fmt.Println("Error writing sum")
		os.Exit(1)
	} else {
		fmt.Printf("Wrote sum %d to %s\n", number, testfile)
	}
}

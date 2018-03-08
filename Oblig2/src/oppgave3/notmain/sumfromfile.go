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
	//reads file
	file, err := ioutil.ReadFile(testfile)
	if err != nil {
		fmt.Printf("\n%s could not be read\n", testfile)
		os.Exit(1)
	}
	//Converts file to a single string
	fileString := string(file)
	//Splits string into a string slice, each position representing a single character
	splitString := []string(strings.Split(fileString, "\n"))

	//int slice with 2 positions
	var ints [2]int

	//Converts all strings into int and puts them in the int slice
	for i := 0; i < len(splitString); i++ {
		ints[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			fmt.Println("Could not convert numbers in file to int")
			os.Exit(1)
		}
	}
	//sums position 0 and 1 of int slice
	total := ints[0] + ints[1]

	return total
}

//writes sum to the file
func WriteSum(number int) {
	//Creates file to be written to, overrides file used previously
	file, err := os.Create(testfile)
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}
	//defer waits for function to finish before file closes
	defer file.Close()

	//Writes number to file
	_, err = file.WriteString(fmt.Sprintf("%d", number))
	if err != nil {
		fmt.Println("Error writing sum")
		os.Exit(1)
	} else {
		fmt.Printf("Wrote sum %d to %s\n", number, testfile)
	}
}

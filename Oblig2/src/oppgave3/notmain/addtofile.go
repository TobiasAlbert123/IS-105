package notmain

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
)

func AddToFile() {
	WriteNumbers(Input())
}

func PrintResult() {
	fmt.Println("Oppgave 3")
	file, err := ioutil.ReadFile("../testfile.txt")
	if err != nil {

	}
	fileString := string(file)
	total, err := strconv.Atoi(fileString)
	if err != nil {

	}
	fmt.Printf("Total is %d\n", total)
}

func WriteNumbers(number1, number2 int) {
	file, err := os.Create("../testfile.txt")
	if err != nil {
		fmt.Println("Error creating file")
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d\n%d" , number1, number2))
	if err != nil {
		fmt.Println("Error writing")
	} else {
		fmt.Printf("Wrote numbers %d and %d to testfile.txt\n", number1, number2)
	}
}

func Input() (int, int) {
	var number1 int
	var number2 int
	fmt.Println("Enter first number")
	_, err := fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number")
	_, err = fmt.Scanf("%d", &number2)
	if err != nil {
	}
	return number1, number2
}

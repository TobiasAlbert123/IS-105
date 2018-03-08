package notmain

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
)

//name of file that all funcs read / write from
var testfile = "readwritenumbers.txt"

//takes input from terminal and returns them
func Input() (int, int) {
	//initialises numbers to be written
	var number1 int
	var number2 int

	//Writes first number
	fmt.Println("Enter first number")
	_ , err := fmt.Scanf("%d", &number1)
	if err != nil {
		fmt.Println("Could not write first number")
		os.Exit(1)
	}

	//Writes second number
	fmt.Println("Enter second number")
	_, err = fmt.Scanf("%d", &number2)
	if err != nil {
		//repeating this line is needed for program to run in Windows cmd for some reason
		_, err = fmt.Scanf("%d", &number2)
		if err != nil {
			fmt.Println("Could not write second number")
			os.Exit(1)
		}
	}

	return number1, number2
}

//writes numbers to file
func WriteNumbers(number1, number2 int) {
	//file should appear in Oblig2 dir if run in Goland
	//file should appear in 3b_run dir if run in Windows cmd
	file, err := os.Create(testfile)
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}
	//defer waits for function to finish before file closes
	defer file.Close()

	//Writes number1 and number2 to file, separated by new line
	_, err = file.WriteString(fmt.Sprintf("%d\n%d" , number1, number2))
	if err != nil {
		fmt.Println("Error writing numbers")
		os.Exit(1)
	}

	fmt.Printf("Wrote numbers %d and %d to %s\n", number1, number2, testfile)
}

//prints sum from file
func PrintResult() {
	fmt.Printf("Printing sum from %s\n", testfile)
	file, err := ioutil.ReadFile(testfile)
	if err != nil {
		fmt.Println("Could not read sum from file")
		os.Exit(1)
	}
	fileString := string(file)
	total, err := strconv.Atoi(fileString)
	if err != nil {
		fmt.Println("Could not convert sum to int")
		os.Exit(1)
	}

	fmt.Printf("Total is %d\n", total)

	//deletes file when done
	//os.Remove(testfile)
}

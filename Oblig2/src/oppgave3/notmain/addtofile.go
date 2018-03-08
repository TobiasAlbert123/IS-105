package notmain

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"time"
)

//name of file that all funcs read / write from
var testfile = "testfile.txt"

//writes numbers to file
func WriteNumbers(number1, number2 int) {
	//file will appear in Oblig2 dir if run in Goland
	//file will appear in 3b_run dir if run in Windows cmd
	file, err := os.Create(testfile)
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d\n%d" , number1, number2))
	if err != nil {
		fmt.Println("Error writing numbers")
		os.Exit(1)
	} else {
		fmt.Printf("Wrote numbers %d and %d to %s\n", number1, number2, testfile)
	}
}

//takes input from terminal and returns them
func Input() (int, int) {
	var number1 int
	var number2 int
	fmt.Println("Enter first number")
	_ , err := fmt.Scanf("%d", &number1)
	if err != nil {
		fmt.Println("Could not write first number")
		os.Exit(1)
	}
	fmt.Println("Enter second number")
	time.Sleep(time.Duration(1)*time.Second)
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
}

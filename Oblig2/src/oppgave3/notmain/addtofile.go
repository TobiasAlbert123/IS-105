package notmain

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"time"
)

var testfile = "testfile.txt"

func AddToFile() {
	WriteNumbers(Input())
}

func PrintResult() {
	fmt.Printf("Printing sum from %s\n", testfile)
	file, err := ioutil.ReadFile(testfile)
	if err != nil {

	}
	fileString := string(file)
	total, err := strconv.Atoi(fileString)
	if err != nil {
		HandleError()
	}
	fmt.Printf("Total is %d\n", total)
}

func WriteNumbers(number1, number2 int) {
	//file will appear in Oblig2 dir if run in Goland
	//file will appear in 3b_run dir if run in Windows cmd
	file, err := os.Create(testfile)
	if err != nil {
		fmt.Println("Error creating file")
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d\n%d" , number1, number2))
	if err != nil {
		fmt.Println("Error writing")
	} else {
		fmt.Printf("Wrote numbers %d and %d to %s\n", number1, number2, testfile)
	}
}

func Input() (int, int) {
	var number1 int
	var number2 int
	fmt.Println("Enter first number")
	_ , err := fmt.Scanf("%d", &number1)
	if err != nil {
		fmt.Println("Could not write first number")
	}
	fmt.Println("Enter second number")
	time.Sleep(time.Duration(1)*time.Second)
	_, err2 := fmt.Scanf("%d", &number2)
	if err2 != nil {
		//this line is needed for program to run in Windows cmd for some reason
		_, _ = fmt.Scanf("%d", &number2)
	}
	return number1, number2
}

func HandleError() {

}

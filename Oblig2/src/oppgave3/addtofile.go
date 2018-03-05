package main

import (
	"fmt"
	"os"
)

func main() {
	WriteNumbers(Input())
}

func WriteNumbers(number1, number2 int) {
	file, err := os.Create("Oblig2/src/oppgave3/testfile.txt")
	if err != nil {
		fmt.Println("Error creating file")
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d\n%d\n" , number1, number2))
	if err != nil {
		fmt.Println("Error writing")
	} else {
		fmt.Printf("Wrote numbers %d and %d to testfile.txt", number1, number2)
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

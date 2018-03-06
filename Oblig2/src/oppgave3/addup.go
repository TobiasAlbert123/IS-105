package main

import (
	"fmt"
)

func main() {
	input()
}

//Funksjon A
func input() {
	var number1 int
	var number2 int
	fmt.Print("Enter first number: ")
	_, err := fmt.Scanf("%d\n", &number1)
	fmt.Print("Enter second number: ")
	_, err = fmt.Scanf("%d\n", &number2)
	if err != nil {
		fmt.Printf("\n\nERROR: %s\nReason: You did not enter 2 valid integers\n", err)
	} else {
		fmt.Printf("Total number = %d", addUp(number1, number2))
	}
}

//Funksjon B
func addUp(tall, tall2 int) int{
	return tall + tall2
}
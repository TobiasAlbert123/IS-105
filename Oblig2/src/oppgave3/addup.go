package main

import (
	"fmt"
)

func main() {
	input()
}

func input() {
	var number1 int
	var number2 int
	fmt.Print("Enter first number: ")
	_, err := fmt.Scanf("%d\n", &number1)
	fmt.Print("Enter second number: ")
	_, err = fmt.Scanf("%d\n", &number2)
	if err != nil {
	}
	fmt.Printf("Total number = %d", addUp(number1, number2))
}

func addUp(tall, tall2 int) int{
	return tall + tall2
}

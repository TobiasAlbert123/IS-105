package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Legg sammen to tall fra terminal og summer i en annen funksjon")
	//fmt.Println(addUp(10, 10))
	addUp(input())
}

func input() (int, int) {
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

func addUp(tall, tall2 int) int{
	return tall + tall2
}

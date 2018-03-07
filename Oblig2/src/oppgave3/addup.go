package main

import (
	"fmt"
	"time"
)

func main() {
	go funcA()
	go funcB()
	//Program sleeps for 20 secs before finishing
	//to prevent program from finishing before funcs are done
	time.Sleep(time.Duration(20)*time.Second)
}

var ch = make(chan int, 2)
var ch2 = make(chan int)

//Funksjon A
func funcA() {
	var number1 int
	var number2 int
	fmt.Print("Enter first number: ")
	_, err := fmt.Scanf("%d\n", &number1)
	ch <- number1
	fmt.Print("Enter second number: ")
	_, err = fmt.Scanf("%d\n", &number2)
	ch <- number2
	if err != nil {
		fmt.Printf("\n\nERROR: %s\nReason: You did not enter 2 valid integers\n", err)
	} else {
		fmt.Printf("Total number = %d", <- ch2)
	}
}

//Funksjon B
func funcB() {
	sum := <- ch + <- ch
	ch2 <- sum
}
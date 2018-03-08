package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

//all os.Exit have 2s delay to allow the SIGINT exit message to be shown
//instead of other error messages

//channel for SIGINT signal
var sigintCh = make(chan os.Signal, 2)

func main() {
	//SIGINT handling
	signal.Notify(sigintCh, os.Interrupt, syscall.SIGINT)
	go stop()

	//go funcs run simultaneously
	go funcA()
	go funcB()

	//Program sleeps for 1 min before finishing
	//to prevent program from finishing before funcs are done
	time.Sleep(time.Duration(1)*time.Minute)
	fmt.Printf("\n\nProgram timed out, please finish the program within 1 minute next time\n")
	time.Sleep(time.Duration(500))
		os.Exit(1)
}

var ch = make(chan int, 2)
var ch2 = make(chan int)

//Funksjon A
func funcA() {
	number1 := 0
	number2 := 0
	fmt.Print("Enter first number: ")
	_, err := fmt.Scanf("%d\n", &number1)

	//program stops if first number has an error
	if err != nil {
		fmt.Printf("first number is invalid, please type an integer\n")
		time.Sleep(2*time.Second)
		os.Exit(1)
	}
	//first number sent to first channel
	ch <- number1
	fmt.Print("Enter second number: ")
	_, err = fmt.Scanf("%d\n", &number2)

	//program stops if second number has an error
	if err != nil {
		fmt.Printf("second number is invalid, please type an integer\n")
		time.Sleep(2*time.Second)
		os.Exit(1)
	}
	//second number sent to first channel
	ch <- number2

	//sum received from second channel
	fmt.Printf("Total number = %d", <-ch2)

	//Program is successful, exits with code 0
	//(program does not exit otherwise until time.Sleep finishes)
	os.Exit(0)

}

//Funksjon B
func funcB() {
	//first and second number received from first channel
	sum := <- ch + <- ch
	//sum sent to second channel
	ch2 <- sum
}

//listens for SIGINT signal from sigintCh
func stop() {
	<- sigintCh
	fmt.Printf("\nCTRL+C stopped the program before it finished\n")
	os.Exit(1)
}
package main

import (
	"../notmain"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

//channel for SIGINT signal
var sigintCh = make(chan os.Signal, 2)


func main() {
	//SIGINT handling
	signal.Notify(sigintCh, os.Interrupt, syscall.SIGINT)
	go ctrlc()

	//addtofile.go takes input and writes it to file
	notmain.WriteNumbers(notmain.Input())

	//sumfromfile.go reads input, sums it and writes it to file
	notmain.WriteSum(notmain.SumFromFile())

	//addtofile.go reads file and prints
	notmain.PrintResult()
}

//listens for SIGINT signal from sigintCh
func ctrlc() {
	<- sigintCh
	fmt.Println("CTRL+C stopped the program before it finished")
	os.Exit(1)
}

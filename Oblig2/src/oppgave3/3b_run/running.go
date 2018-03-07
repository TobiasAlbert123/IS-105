package main

import (
	"../notmain"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

var ch = make(chan os.Signal, 2)


func main() {
	signal.Notify(ch, os.Interrupt, syscall.SIGINT)
	go ctrlc()

	//addtofile.go takes input and writes it to file
	notmain.AddToFile()
	//sumfromfile.go reads input, sums it and writes it to file
	notmain.WriteSum(notmain.SumFromFile())
	//addtofile.go reads file and prints
	notmain.PrintResult()
}

func ctrlc() {
	<- ch
	fmt.Println("CTRL+C stopped the program before it finished")
	os.Exit(1)
}

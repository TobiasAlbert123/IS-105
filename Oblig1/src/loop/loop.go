package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"time" // or "runtime"
)

func cleanup() {
	fmt.Println("You have interrupted my loop. I will close.")
}

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for i := 0;;i++{
		fmt.Println(i)
	}
}
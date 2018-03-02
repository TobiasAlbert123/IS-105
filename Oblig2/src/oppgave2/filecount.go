package main

import (
	"fmt"
	"os"
	"bufio"
)


//to run in cmd:
//1. Have same map architecture as git
//2. cd [repository path]/oppgave2
//3. go run filecount.go text.txt
func main() {
	//takes argument from cmd line
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Argument missing")
		fmt.Println("Type 'go run fileinfo.go text.txt'")
	}

	fileName := args[1]
	countLines(fileName)

	makeMap()
	somevalues()
	printtest()
	fmt.Println(RuneMap)
}

var srcFolder = "../files/"

//counts lines in file
func countLines(input string) {
	file, _ := os.Open(srcFolder + input)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Printf("'text.txt' has %d lines", lineCount)
}

//similiar to hashmap in java
var RuneMap = make(map[rune]int)

//adds values to map
func makeMap() {
	for i := 0; i < 256; i++ {
		RuneMap[rune(i)] = 0
	}
}

//adds 1 to runes 98 and 110 in map
func somevalues() {
	RuneMap[rune(98)]++
	RuneMap[rune(110)]++
}

//prints runes of value i and their numbers recorded
func printtest() {
	for i := 0; i < 256; i++ {
		fmt.Println(string(i), RuneMap[rune(i)])
	}
}
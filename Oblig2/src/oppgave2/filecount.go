package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
)


//to run in cmd:
//1. Have same map architecture as git
//2. cd [repository path]/oppgave2
//3. go run filecount.go text.txt
func main() {
	countLines(getFile())
	MapMaker(getFile())
	//fmt.Println(RuneMap)
	sortMap()
}

//Gets filename from run argument and adds filepath to the files folder
func getFile() string{
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Argument missing")
		fmt.Println("Type 'go run fileinfo.go text.txt'")
	}
	fileName := args[1]
	srcFolder := "../files/"
	filePath := srcFolder + fileName
	return filePath
}

//counts lines in file
func countLines(filePath string) {
	file, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Printf("'text.txt' has %d lines\n", lineCount)
}

//initialises map containing runes and int
var RuneMap = make(map[rune]int)


//Creates map and feeds values from file into addToMap
func MapMaker(text string) {
	file, err := ioutil.ReadFile(text)
	if err != nil {

	}
	fileString := string(file)
	splitString := []string(strings.Split(fileString, ""))

	for i:= 0; i < len(splitString); i++ {
		addToMap(splitString[i])
	}
}

//slice of all found letters, used for easier comparisons
var foundLetters []string

//adds values to map
func addToMap(singleLetter string) {
	foundIt := false
	for i:= 0; i < len(foundLetters); i++ {
		if singleLetter == foundLetters[i] {
			foundIt = true
		}
	}

	//Converts string to rune to int for easier adding to RuneMap
	runeOfString := []rune(singleLetter)
	intOfRune := int(runeOfString[0])

	//adds 1 to the rune's value in the map if found, else creates new entry with value 1
	if foundIt == true {
		RuneMap[rune(intOfRune)]++
	} else {
		RuneMap[rune(intOfRune)] = 1
		foundLetters = append(foundLetters, singleLetter)
	}
}

func sortMap() {
	for i := 1; i < 6; i++ {
		number := i
		highestCount := 0
		mostUsed := ""
		for i := 0; i < len(foundLetters); i++ {
			runeOfString := []rune(foundLetters[i])
			mentions := RuneMap[runeOfString[0]]
			if mentions > highestCount {
				highestCount = mentions
				mostUsed = foundLetters[i]
			}
		}
		mostUsedRune := []rune(mostUsed)
		fmt.Printf("%d. Rune: '%s' , Counts: %d\n", number, mostUsed, highestCount)
		delete(RuneMap, mostUsedRune[0])
	}
}
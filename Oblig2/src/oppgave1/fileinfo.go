package main

import (
	"fmt"
	"os"
	"log"
)

//to run in cmd:
//1. Have same map architecture as git
//2. cd [repository path]/oppgave1
//3. go run fileinfo.go text.txt
func main () {
	//takes argument from cmd line
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Argument missing")
		fmt.Println("Type 'go run fileinfo.go text.txt'")
	}

	fileName := args[1]

	fmt.Printf("Info for file '%s':\n", fileName)
	fileSize(fileName)
	fileInfo(fileName)
}

//Relative path: Oblig2/src/files
var srcFolder = "../files/"

func fileSize(input string) {
	//finds file with name input in set folder
	path := srcFolder + input
	file, err := os.Stat(path)
		if err != nil {
			ErrorHandling(err)
		}

	bytes := file.Size()
	//bytes as float (for decimal numbers)
	fBytes := float64(bytes)

	//Prints bytes, KB, MB & GB
	fmt.Printf("The file has size:\n")
	fmt.Printf("%d	bytes\n", bytes)
	fmt.Printf("%.2f	KB \n%.2f	MB \n%.2f	GB \n\n", fBytes / (1024), fBytes / (1024 * 1024), fBytes / (1024 * 1024* 1024) )

}

func fileInfo(input string) {
	//finds file with name input in set folder
	path := srcFolder + input
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	mode := fi.Mode()

	fmt.Println("The file:")

	//Check if directory
	if mode.IsDir() {
		fmt.Println("is a directory")
	} else {
		fmt.Println("is not a directory")
	}

	//Check if regular file
	if mode.IsRegular() {
		fmt.Println("is a regular file")
	} else {
		fmt.Println("is not a regular file")
	}

	//Check if file has UNIX permission bits
	fmt.Printf("has UNIX permission bits %s\n", mode.Perm())

	//Check if append-only
	if mode&os.ModeAppend != 0 {
		fmt.Println("is append-only")
	} else {
		fmt.Println("is not append-only")
	}

	//Check if device file
	if mode&os.ModeDevice != 0 {
		fmt.Println("is a device file")
	} else {
		fmt.Println("is not a device file")
	}
}

//Logs error (somewhere?)
func ErrorHandling(input error) {
	log.Fatal(input)
}
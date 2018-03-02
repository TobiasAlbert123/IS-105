package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main2(input string) {
	fmt.Println("Programmet skal returnere informasjon om en fil")
	fileSize(input)
	Info(input)
}

func printAnInput () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}
func main () {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	main2(arg)
}

func fileSize(input string) {
	path := "Oblig2/src/files/" + input
	fi, err := os.Stat(path)
	if err != nil {

	}
	bytes := fi.Size()
	kiloBytes := fi.Size()/1024
	megaBytes := kiloBytes/1024
	gigaBytes := megaBytes/1024
	fmt.Printf("The file has size (rounded down):\n%d bytes\n%d KB\n%d MB\n%d GB\n\n", bytes, kiloBytes, megaBytes, gigaBytes)

}

func Info(input string) {
	path := "Oblig2/src/files/" + input
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The file")

	mode := fi.Mode()

	if mode.IsDir() {
		fmt.Println("is a directory")
	} else {
		fmt.Println("is not a directory")
	}

	if mode.IsRegular() {
		fmt.Println("is a regular file")
	} else {
		fmt.Println("is not a regular file")
	}

	// Has Unix permission bits: -rwxrwxrwx


	if mode&os.ModeAppend != 0 {
		fmt.Println("is append-only")
	} else {
		fmt.Println("is not append-only")
	}

	if mode&os.ModeDevice != 0 {
		fmt.Println("is a device file")
	} else {
		fmt.Println("is not a device file")
	}

	// Is/Is not a Unix character device

	// Is/Is not a Unix block device

	if mode&os.ModeSymlink != 0 {
		fmt.Println("is a symbolic link")
	} else {
		fmt.Println("is not a symbolic link")
	}

	if mode&os.ModeNamedPipe != 0 {
		fmt.Println("is a named pipe")
	} else {
		fmt.Println("is not a named pipe")
	}
}
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Oppgave 3")
	file, err := os.Open("Oblig2/src/oppgave3/testfile.txt")
	if err != nil {

	}
	buf := make([]byte, 1024)
	fmt.Println(file.Read(buf))
}

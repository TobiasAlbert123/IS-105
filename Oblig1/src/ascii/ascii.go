package main

import (
	"fmt"
	"strconv"
)

//int slice representing " € ÷ ¾ dollar "
var letters = [16]int{34,32,8364,32,247,32,190,32,100,111,108,108,97,114,32,34}

func main() {
	for i:=0x80; i <= 0xFF; i++ {
		IterateOverASCIIStringLiteral(i)
	}
	ExtendedASCIIText(letters)
}

func IterateOverASCIIStringLiteral(i int) {
	//Converts to hex
	hex := fmt.Sprintf("%X", i)
	//Converts to string (ASCII character)
	str := string(i)
	//Converts to int64 and then binary
	n := int64(i)
	bin := strconv.FormatInt(n, 2)
	//Prints hex, ASCII, binary
	fmt.Printf("%s %s %s\n", hex, str, bin)
}

func ExtendedASCIIText(letters [16]int) {
	//adds letters to  allLetters based on their ASCII value listed in letters
	allLetters := string("")
	for i:= 0; i < len(letters); i++ {
		allLetters += string(letters[i])
	}
	fmt.Printf("%s\n", allLetters)
}
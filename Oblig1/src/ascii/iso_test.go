package main

import (
	"fmt"
	//"strings"
	"testing"
)

func TestExtendedASCIIText(t *testing.T) {
	//letters is input for ExtendedASCIIText
	for i := 0; i<len(letters); i++ {
		character := string(letters[i])
		number := letters[i]
		if number >= 0x80 && number <= 0xFF {
			fmt.Printf("'%s' is in the extended ASCII table (pos %d)\n", character, number)
		} else if letters[i] >= 0x00 && number < 0x80{
			t.Fail()
			fmt.Printf("'%s' is in the normal ASCII table (pos %d)\n", character, number)
		} else {
		t.Fail()
		fmt.Printf("'%s' is NOT in the extended ASCII table (pos %d)\n", character, number)
		}
	}
}

/*
* Attempt at 4C
 */
/*func ISO(input string) {
	//splits input into a string array
	splitString := []string(strings.Split(input, ""))
	//takes each character and converts to rune, then converts to int
	for i := 0; i < len(splitString); i++ {
		character := splitString[i]
		CharToRune := []rune(character)
		RuneToInt := int(CharToRune[0])
		//prints whether its part of ASCII extended or not
		if RuneToInt >= 0x80 && RuneToInt <= 0xFF {
			fmt.Printf("%s is in extended ASCII table at position %d\n", character, RuneToInt)
		} else {
			fmt.Printf("%s is NOT in extended ASCII table (position %d)\n", character, RuneToInt)
		}
	}
}*/

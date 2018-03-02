package main

import "fmt"

func main() {
	fmt.Println("Programmet skal lese en tekst-fil, skrive ut totalt antall linjer og antall for fem runes som forekommer hyppigst i filen")
	makeMap()
	somevalues()
	printtest()
	fmt.Println(RuneMap)
	}

var RuneMap = make(map[rune]int)

func makeMap() {
	for i := 0; i < 256; i++ {
		RuneMap[rune(i)] = 0
	}
}

func somevalues() {
	RuneMap[rune(98)]++
	RuneMap[rune(110)]++
}

func printtest() {
	for i := 0; i < 256; i++ {
		fmt.Println(string(i), RuneMap[rune(i)])
	}
}
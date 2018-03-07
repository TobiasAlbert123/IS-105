package main

import (
	"../notmain"
)

func main() {
	//addtofile.go takes input and writes it to file
	notmain.AddToFile()
	//sumfromfile.go reads input, sums it and writes it to file
	notmain.WriteSum(notmain.SumFromFile())
	//addtofile.go reads file and prints
	notmain.PrintResult()
}

package main

import (
	"../notmain"
)

func main() {
	notmain.AddToFile()
	notmain.WriteSum(notmain.SumFromFile())
	notmain.PrintResult()
}

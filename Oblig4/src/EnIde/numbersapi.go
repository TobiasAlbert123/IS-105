package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"strconv"
)

func main () {
	borings := []int{}
	unremarkables := []int{}
	uninterestings := []int{}
	for i := 2100; i < 2160; i++ {
		Response, err := http.Get(makeUrl(i))
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		Body, err := ioutil.ReadAll(Response.Body)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		info := string(Body)
		num := strconv.Itoa(i)
		switch info {
		case num + " is a boring number.":
			borings = append(borings, i)
		case num + " is an unremarkable number.":
			unremarkables = append(unremarkables, i)
		case num + " is an uninteresting number.":
			uninterestings = append(uninterestings, i)
		default: fmt.Println(info)
		}
	}
	fmt.Printf("Boring numbers: %d\n", len(borings))
	fmt.Printf("Unremarkable numbers: %d\n", len(unremarkables))
	fmt.Printf("Uninteresting numbers: %d\n", len(uninterestings))
}

func makeUrl(number int) string{
	base1 := "http://numbersapi.com/"
	base2 := "/math"
	url := base1 + strconv.Itoa(number) + base2
	return url
}

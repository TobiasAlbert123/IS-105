package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
	"strconv"
)

type Items struct {
	Items	[]Item
}

type Item struct {
	Name	string
	Price	int
}


func main() {
	url1 := "https://www.komplett.no/product/950622/datautstyr/pc-komponenter/prosessorer/intel-core-i9-7980xe-prosessor"
	url2 := "https://www.komplett.no/product/927754/mobil/mobiltelefon/apple-iphone-6-32gb-graa"
	urls := []string{
		url1,
		url2,
		"https://www.komplett.no/product/903116/datautstyr/lagring/harddiskerssd/ssd-m2/samsung-960-evo-500gb-m2-pcie-ssd",
		"https://www.komplett.no/product/985604/datautstyr/pc-komponenter/skjermkort/nvidia-titan-v",
	}
	names := []string{}
	prices := []int{}
	for i := 0; i < len(urls); i++ {
		root := getPage(urls[i])
		names = append(names, formatName(get2Thing("visible: addToCart.buttonClicked", root)))
		fmt.Printf("Navn: %s\n", formatName(get2Thing("visible: addToCart.buttonClicked", root)))
		prices = append(prices, formatPrice(getThing("price", root)))
		fmt.Printf("Pris: %d %s\n\n", formatPrice(getThing("price", root)), getThing("priceCurrency", root))
	}
	fmt.Println("AGAIN")
	for i := 0; i < len(names); i++ {
		fmt.Printf("Navn: %s, Pris: %d\n\n", names[i], prices[i])
	}

	//does not work
	listed := sortByHighestPrice(prices)
	for i := 0; i < len(listed); i++ {
		fmt.Printf("Number %d price: %d\n", i, listed[i])
	}
}

//does not work
func sortByHighestPrice(input []int) []int{
	temp := input
	highest := 0
	listed := []int{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(temp); j++ {
			if temp[j] > highest {
				highest = temp[j]
			}
		}
		listed = append(listed, highest)
		temp = append(temp, highest)
	}
	return listed
}

func formatName(input string) string {
	toRemove := "Lagt til "
	input = strings.TrimPrefix(input, toRemove)
	return input
}

func formatPrice(input string) int {
	number, _ := strconv.Atoi(strings.TrimSuffix(input, ".00"))
	return number
}


func getPage(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return root
}

func getElementByItemProp(itemprop string, n *html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "itemprop" && a.Val == itemprop {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = getElementByItemProp(itemprop, c); ok {
			return
		}
	}
	return
}

func getElementByDataBind(dataBind string, n *html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "data-bind" && a.Val == dataBind {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = getElementByDataBind(dataBind, c); ok {
			return
		}
	}
	return
}

func get2Thing(firstKey string, root *html.Node) string{
	element, ok := getElementByDataBind(firstKey, root)
	if !ok {
		log.Fatal("element not found")
	}
	secondKey := "aria-label"
	for _, a := range element.Attr {
		if a.Key == secondKey {
			//fmt.Print(a.Val)
			return a.Val
		}
	}
	log.Fatalf("element missing %s", secondKey)
	return "Error"
}

func getThing(firstKey string, root *html.Node) string{
	element, ok := getElementByItemProp(firstKey, root)
	if !ok {
		log.Fatal("element not found")
	}
	secondKey := "content"
	for _, a := range element.Attr {
		if a.Key == secondKey {
			//fmt.Print(a.Val)
			return a.Val
		}
	}
	log.Fatalf("element missing %s", secondKey)
	return "Error"
}
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func main() {
	url1 := "https://www.komplett.no/product/950622/datautstyr/pc-komponenter/prosessorer/intel-core-i9-7980xe-prosessor"
	url2 := "https://www.komplett.no/product/927754/mobil/mobiltelefon/apple-iphone-6-32gb-graa"
	urls := []string{
		url1,
		url2,
		"https://www.komplett.no/product/903116/datautstyr/lagring/harddiskerssd/ssd-m2/samsung-960-evo-500gb-m2-pcie-ssd",
	}
	for i := 0; i < len(urls); i++ {
		root := getPage(urls[i])
		fmt.Printf("Navn: %s\n", formatName(get2Thing("visible: addToCart.buttonClicked", root)))
		fmt.Printf("Pris: %s%s\n\n", getThing("price", root), getThing("priceCurrency", root))
	}
}

func formatName(input string) string {
	toRemove := "Lagt til "
	input = strings.Replace(input, toRemove, "", 1)
	return input
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
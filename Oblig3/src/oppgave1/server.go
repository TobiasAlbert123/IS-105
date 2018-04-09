package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", helloClient)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloClient (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Hello, client.")
}
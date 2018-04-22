package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"bufio"
	"time"
)

func main() {
	fmt.Printf("Program starting...\n\n")
	//creates UDP and TCP connections independently as 'go' functions
	go connUDP(udpQuote())
	go connTCP(tcpQuote())

	//prevents program from instantly exiting
	time.Sleep(time.Minute * 10)
	fmt.Println("Program finished. Exiting...")
}


//type structs for getting quotes from a json file
type Contents struct {
	Contents	Quotes	`json:"contents"`
}

type Quotes struct {
	Quotes	[]Quote `json:"quotes"`
}

type Quote struct {
	Quote	string	`json:"quote"`
}

var localAddress = ":17"

//simple error handling
func handleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

//UDP connection
func connUDP(quote string) {
	serverAddress, err := net.ResolveUDPAddr("udp", localAddress)
	handleError(err)

	serverConnection, err := net.ListenUDP("udp", serverAddress)
	handleError(err)
	defer serverConnection.Close()

	buffer := make([]byte, 1024)

	for {
		_ , address, err := serverConnection.ReadFromUDP(buffer)
		handleError(err)
		//fmt.Println("Received: ", string(buffer[0:length]), " from ", address)

		message := "UDP quote: "

		serverConnection.WriteToUDP([]byte(message + quote + "\n\n"), address)
	}
}

//TCP connection
func connTCP(quote string) {
	listen, err := net.Listen("tcp", localAddress)
	handleError(err)

	connection, err := listen.Accept()
	handleError(err)

	for {
		message, _ := bufio.NewReader(connection).ReadString('\n')
		//fmt.Print("Message received: ", string(message))
		message = "TCP quote: "
		connection.Write([]byte(message + quote + "\n\n"))
	}
}

//opens and unmarshals local file with quotes
func openQuotes() Contents {
	file, err := os.Open("quotes.json")
	handleError(err)
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	quotes := Contents{}

	json.Unmarshal(bytes, &quotes)

	return quotes

}

//gets the quote for UDP response
func udpQuote() string{
	quote := openQuotes().Contents.Quotes[0].Quote
	fmt.Println(quote)
	return quote
}

//gets the quote for TCP response
func tcpQuote() string {
	quote := openQuotes().Contents.Quotes[1].Quote
	fmt.Println(quote)
	return quote
}
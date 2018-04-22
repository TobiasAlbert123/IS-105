package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"bufio"
	"strings"
	"time"
)

func main() {
	fmt.Println("Program starting")
	go connUDP(udpQuote())
	go connTCP(tcpQuote())
	time.Sleep(time.Minute * 10)
	fmt.Println("Program finished. Exiting...")
}

type Contents struct {
	Contents	Quotes	`json:"contents"`
}

type Quotes struct {
	Quotes	[]Quote `json:"quotes"`
}

type Quote struct {
	Quote	string	`json:"quote"`
}

func LolErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func connUDP(quote string) {
	serverAddress, err := net.ResolveUDPAddr("udp", ":8081")
	LolErr(err)

	serverConnection, err := net.ListenUDP("udp", serverAddress)
	LolErr(err)
	defer serverConnection.Close()

	buffer := make([]byte, 1024)

	for {
		n, address, err := serverConnection.ReadFromUDP(buffer)
		fmt.Println("Received: ", string(buffer[0:n]), " from ", address)
		LolErr(err)

		serverConnection.WriteToUDP([]byte(quote+"\n"), address)
	}
}

func connTCP(quote string) {
	listen, err := net.Listen("tcp", ":8081")
	LolErr(err)

	connection, err := listen.Accept()
	LolErr(err)

	for {
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("Message received: ", string(message))
		newmessage := strings.ToUpper(message)
		connection.Write([]byte(quote + newmessage + "\n"))
	}
}

func openQuotes() Contents {
	file, err := os.Open("quotes.json")
	LolErr(err)
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	q := Contents{}

	json.Unmarshal(bytes, &q)

	return q

}

func tcpQuote() string {
	page := openQuotes()
	quote := page.Contents.Quotes[0].Quote
	fmt.Println(quote)
	return quote
}

func udpQuote() string{
	page := openQuotes()
	quote := page.Contents.Quotes[1].Quote
	fmt.Println(quote)
	return quote
}
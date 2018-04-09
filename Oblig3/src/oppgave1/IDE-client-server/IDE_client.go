package IDE_client_server

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func main() {

	//tries to open a tcp connection at port 8080
	conn, _ := net.Dial("tcp", "localhost:8080")

	//reads a new input line
	scanner := bufio.NewReader(os.Stdin)

	//shutdown with ctrl+c
	for {

		fmt.Print("Send text: ")

		//reads input
		text, _ := scanner.ReadString('\n')

		//sends input over connection
		fmt.Fprintf(conn, text + "\n")


		message, _:= bufio.NewReader(conn).ReadString('\n')


		fmt.Print("Message from server: " + message)
	}

}
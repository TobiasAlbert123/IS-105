package IDE_client_server

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"net/http"
	"time"
)

func main() {

	//listens at port 8080 for a tcp connection
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error: %v \nPlease try again", err)
		return
	}

	//accepts a connection
	conn, _ := listen.Accept()


	http.Handle("/lol", hello())

	//shutdown with ctrl+c
	for {
		//Reads a line from connection
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message received from client: ", string(message))

		//new message to be sent
		newMessage := strings.ToUpper(message)

		newMessage = "lame-o " + newMessage

		//sends new message back to client
		conn.Write([]byte(newMessage + "\n"))

		fmt.Printf("Message sent to client: %s\n", newMessage)

		//Prevents loop from running at max speed when there isnt a client
		time.Sleep(time.Second*2)
	}
}

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "Hello CLIENT")
		fmt.Println("itdidit")

	})
}
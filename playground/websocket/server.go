package websocket

import (
	"fmt"
	"net"
	"log"
	"time"
)

func checkError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func handleClientConnection(conn net.Conn) {
	for {
		data := "Hello World"
		conn.Write([]byte(data))
		time.Sleep(100*time.Millisecond)
	}
	conn.Close()
}

func Run() {
	// create a socket
	// bind a socket
	// listen to that socket
	lis, err := net.Listen("tcp", ":8080")

	checkError(err)

	// close socket
	defer lis.Close()

	// event loop
	// accept client connections
	// close client connections

	for {
		client, err := lis.Accept()

		checkError(err)

		fmt.Println("New Connection")

		go handleClientConnection(client)
	}


}
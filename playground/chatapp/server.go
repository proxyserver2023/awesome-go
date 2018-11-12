package chatapp

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	checkError(err)
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}

}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				delete(clients, client)
				client.Close()
				log.Printf("error: %v", err)
			}
		}
	}
}

func Run() {
	// listener, err := net.Listen("tcp", ":8080")
	// checkError(err)
	// defer listener.Close()

	// create a file server
	fs := http.FileServer(http.Dir("/home/alamin"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	log.Println("Http Server started on :8080")
	err := http.ListenAndServe(":8080", nil)

	checkError(err)

	go handleMessages()

	// for {
	// 	client, err := listener.Accept()
	// 	checkError(err)
	// 	fmt.Println("New Connection")
	// 	go handleChatConnection(client)
	// }

}

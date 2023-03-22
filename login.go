package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var connections = make(map[*websocket.Conn]bool)
var broadcast = make(chan message)

func handleWebSocket(conn *websocket.Conn) {
	// add connection to list of active connections
	connections[conn] = true

	for {
		// read message from connection
		var msg message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading message:", err)
			delete(connections, conn)
			return
		}

		// broadcast message to all other connections
		broadcast <- msg
	}
}

func main() {
	go func() {
		for {
			// wait for incoming message to broadcast
			msg := <-broadcast

			// send message to all active connections
			for conn := range connections {
				err := conn.WriteJSON(msg)
				if err != nil {
					fmt.Println("Error broadcasting message:", err)
					delete(connections, conn)
				}
			}
		}
	}()
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received: %s\n", p)

		// 受信したメッセージをエコーバック
		if err := ws.WriteMessage(messageType, p); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	fmt.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

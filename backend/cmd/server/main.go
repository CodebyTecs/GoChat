package server

import (
	"log"
	"net/http"

	"GoChat/internal/server/websocket"
)

func main() {
	http.HandleFunc("/ws", websocket.HandleWebSocket)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

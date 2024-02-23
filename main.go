package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadhabibi14/learn-go-websocket/pkg/socket"
)

// define our websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := socket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go socket.Writer(ws)
	socket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple server")
	})

	http.HandleFunc(`/ws`, serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}

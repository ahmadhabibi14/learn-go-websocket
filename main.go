package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadhabibi14/learn-go-websocket/pkg/socket"
)

func serveWs(pool *socket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket endpoint hit")
	conn, err := socket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &socket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := socket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}

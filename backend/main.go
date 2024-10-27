package main

import (
	customwebsocket "chatapplication/websocket"
	"fmt"
	"net/http"
)

func serverWs(pool *customwebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("server with socket runs and works")
	conn, err := customwebsocket.Upgrader(w, r)
	if err != nil {
		fmt.Println("err with serverWs ", err)
		return
	}

	client := &customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()

}

func setupRoutes() {
	fmt.Println("this is working")
	pool := customwebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})

}

func main() {
	go setupRoutes()

	http.ListenAndServe(":9000", nil)

}

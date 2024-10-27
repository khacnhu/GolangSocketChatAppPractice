package customwebsocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, msg, err := c.Conn.ReadMessage()

		if err != nil {
			fmt.Println("err in Read ", err)
			return
		}

		m := Message{Type: msgType, Body: string(msg)}

		c.Pool.Broadcast <- m
		fmt.Println("msg received ====> \n", m)

	}

}

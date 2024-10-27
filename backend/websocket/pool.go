package customwebsocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("client connect websocket ", client)
			fmt.Println("totle connect pool: ", len(pool.Clients))
			for kjoin, _ := range pool.Clients {
				fmt.Println("kjoin = ", kjoin)
				kjoin.Conn.WriteJSON(Message{Type: 1, Body: "New User join room"})
			}
			break

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			for kdelete, _ := range pool.Clients {
				fmt.Println("kdelete = ", kdelete)
				kdelete.Conn.WriteJSON(Message{Type: 1, Body: "New User disconnect room"})
			}
			break
		case msg := <-pool.Broadcast:
			fmt.Println("broadcasting a message")
			for kbroad, _ := range pool.Clients {
				fmt.Println("kbroad = ", kbroad)

				if err := kbroad.Conn.WriteJSON(msg); err != nil {
					fmt.Println("err ", err)
					return

				}

			}

		}
	}
}

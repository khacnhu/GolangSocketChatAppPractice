package customwebsocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Websocket connect error: ", err)
		return nil, err
	}

	return conn, nil
}

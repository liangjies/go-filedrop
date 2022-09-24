package global

import "github.com/gorilla/websocket"

var (
	TextTemp []byte = []byte("Hello World")
	WSConn   []*websocket.Conn
)

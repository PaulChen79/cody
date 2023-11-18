package domain

import "github.com/gorilla/websocket"

type CodeLanguageType string

var (
	Golang    CodeLanguageType = "golang"
	WsClients                  = make(map[uint]*websocket.Conn) // IDE id to conn
)

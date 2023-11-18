package domain

import "github.com/gorilla/websocket"

type Service interface {
	HandleRunCodeWsMessage(ws *websocket.Conn)
}

package websocket

import (
	"cody/domain"
	"cody/internal/provider"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"go.uber.org/zap"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func RunCode(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.S().Warnf("RunCode - err: %v", err)
		return
	}

	domain.WsClients[1] = ws

	defer func() {
		ws.Close()
		delete(domain.WsClients, 1)
	}()

	svc, err := provider.NewService()
	if err != nil {
		zap.S().Warnf("RunCode - err: %v", err)
		return
	}

	svc.HandleRunCodeWsMessage(ws)
}

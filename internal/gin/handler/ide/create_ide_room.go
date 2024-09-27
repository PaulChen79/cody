package ide

import (
	"bytes"
	"cody/domain"
	"cody/internal/gin/handler"
	"cody/internal/provider"
	"io"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateIdeChatRoomRequest struct {
	CreatorID  uint  `json:"creator_id" binding:"required"`
	ChatRoomID *uint `json:"chat_room_id"`
}

func CreateIdeChatRoom(c *gin.Context) {
	var req CreateIdeChatRoomRequest
	reqBody, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(reqBody))
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("CreateIdeChatRoom - err: %v, req: %v", err, string(reqBody))
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	svc, err := provider.NewService()
	if err != nil {
		zap.S().Infof("CreateIdeChatRoom - err: %v", err)
		handler.Failed(c, domain.ErrorServer, err.Error())
		return
	}

	ideChatRooms, err := svc.CreateIdeChatRoom(req.CreatorID, req.ChatRoomID)
	if err != nil {
		zap.S().Infof("CreateIdeChatRoom - err: %v", err)
		handler.Failed(c, domain.ErrorServer, err.Error())
		return
	}

	handler.Success(c, ideChatRooms)
}

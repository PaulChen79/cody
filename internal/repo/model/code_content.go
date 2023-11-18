package gorm

import "cody/domain"

type CodeContent struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	IdeChatRoomID uint   `gorm:"index"`
	Code          []byte `gorm:"type:bytea;not null"`
	FileName      string `gorm:"type:varchar(150);not null"`
}

func CodeContentModelToDomain(codeContent *CodeContent) *domain.CodeContent {
	return &domain.CodeContent{
		ID:            codeContent.ID,
		IdeChatRoomID: codeContent.IdeChatRoomID,
		Code:          codeContent.Code,
		FileName:      codeContent.FileName,
	}
}

func CodeContentDomainToModel(codeContent *domain.CodeContent) *CodeContent {
	return &CodeContent{
		ID:            codeContent.ID,
		IdeChatRoomID: codeContent.IdeChatRoomID,
		Code:          codeContent.Code,
		FileName:      codeContent.FileName,
	}
}

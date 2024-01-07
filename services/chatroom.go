package services

import (
	"github.com/gin-gonic/gin"
	"go_im/lib/mysqllib"
	"go_im/models"
)

func CreateChatroomMessage(c *gin.Context, info models.Chatroom) error {
	info.Ip = c.ClientIP()
	return mysqllib.Client.Create(&info).Error
}

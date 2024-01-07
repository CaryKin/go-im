/*
*

	@author: jixiaogang
	@since: 2024/1/5
	@desc: //TODO

*
*/
package routers

import (
	"github.com/gin-gonic/gin"
	"go_im/controllers/chatroom"
)

func Init(router *gin.Engine) {

	chatroomRouter := router.Group("/chatroom")
	{
		chatroomRouter.GET("/getMessage", chatroom.GetMessage)
		chatroomRouter.POST("/sendMessage", chatroom.SendMessage)
	}
}

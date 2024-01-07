package chatroom

import (
	"github.com/gin-gonic/gin"
	"go_im/common/cache"
	"go_im/common/code"
	"go_im/models"
	"go_im/models/response"
	"go_im/services"
)

// GetMessage 给全员发送消息
func GetMessage(c *gin.Context) {
	// 获取参数
	appIdStr := c.PostForm("appId")

	data := make(map[string]interface{})
	messages, err := cache.GetChatroomMessage(appIdStr)
	if err != nil {
		data["sendResultsErr"] = err.Error()
	}
	data["data"] = messages

	response.SuccessData(c, data)
}

// SendMessage 发送消息
func SendMessage(c *gin.Context) {
	var info models.Chatroom
	if err := c.ShouldBind(&info); err != nil {
		response.ErrorMessage(c, code.ParameterIllegal, "参数验证失败")
	}

	if err := services.CreateChatroomMessage(c, info); err != nil {
		response.Error(c, code.OperationFailure)
	} else {
		response.Success(c)
	}
}

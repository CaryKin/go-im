/*
*

	@author: jixiaogang
	@since: 2024/1/2
	@desc: //TODO

*
*/
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go_im/lib/redislib"
	"go_im/models"
)

const (
	chatroomMessageKey = "acc:hash:chatroom:message:%s" // 全部的服务器
)

func getChatroomMessageKey(roomId string) (key string) {
	key = fmt.Sprintf("%s%s", chatroomMessageKey, roomId)
	return
}

// GetChatroomMessage 获取聊天室消息
func GetChatroomMessage(roomId string) (messages []*models.Chatroom, err error) {
	key := getChatroomMessageKey(roomId)
	redisClient := redislib.GetClient()
	messageMap, err := redisClient.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		fmt.Println("getChatroomMessage", key, err)

		return
	}

	for _, value := range messageMap {
		message := &models.Chatroom{}
		err = json.Unmarshal([]byte(value), message)
		if err != nil {
			fmt.Println("获取直播间消息 json Unmarshal", roomId, err)
			return
		}
		messages = append(messages, message)
	}
	return
}

// SetChatroomMessage 设置直播间消息
func SetChatroomMessage(roomId string, message *models.Chatroom) {

	key := getChatroomMessageKey(roomId)
	redisClient := redislib.GetClient()
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("设置直播间消息转换json失败:", err)
	}
	err = redisClient.LPush(context.Background(), key, data).Err()
	if err != nil {
		fmt.Println("存储消息失败", roomId, err)
		return
	}
	redisClient.LTrim(context.Background(), key, 0, 49)
}

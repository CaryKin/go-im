package models

type Chatroom struct {
	MsgId   string `json:"msgId" form:"msgId" gorm:"column:msg_id"`
	RoomId  uint32 `json:"roomId" form:"roomId" gorm:"column:room_id"`
	Message string `json:"message" form:"message" gorm:"column:message"`
	MsgType uint32 `json:"msgType" form:"msgType" gorm:"column:msg_type"`
	Ext     string `json:"ext" form:"ext" gorm:"column:ext"`
	UserId  string `json:"userId" form:"userId" gorm:"column:user_id"`
	Ip      string `json:"ip" gorm:"column:ip"`
	BaseStruct
}

func (Chatroom) TableName() string {
	return "message_chatroom"
}

func SetChatroomMessage(roomId uint32, content string, userId string, msgId string) (message *Chatroom) {
	message = &Chatroom{
		MsgId:   msgId,
		RoomId:  roomId,
		Message: content,
		UserId:  userId,
		MsgType: 1,
	}
	return
}

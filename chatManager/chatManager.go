package chatManager

import (
	"encoding/json"
	"liansushe/dao"
	"log"

	"github.com/gorilla/websocket"
)

// 表示两个人的会话
type chatGroup struct {
	Conns   map[string]*websocket.Conn
	Message []messageT
}

type messageT struct {
	UserID  string
	Message string
}

type messageToSend struct {
	UserID   string
	UserName string
	Message  string
	IsSelf   bool
}

var chatGroups = make(map[string]*chatGroup)

func Register(userID string, targetUserID string, conn *websocket.Conn) error {
	log.Println("[Register]", userID, targetUserID)
	chatGroupID := getChatGroupID(userID, targetUserID)
	if _, ok := chatGroups[chatGroupID]; ok {
		chatGroups[chatGroupID].Conns[userID] = conn
	} else {
		chatGroups[chatGroupID] = &chatGroup{
			Conns:   map[string]*websocket.Conn{userID: conn},
			Message: make([]messageT, 0),
		}

	}
	log.Println("all conns：")
	for id := range chatGroups[chatGroupID].Conns {
		log.Println(id)
	}

	if length := len(chatGroups[chatGroupID].Message); length > 0 {
		for i := max(0, length-5); i < length; i++ {
			// uid 是发这条消息的人的id，如果uid等于注册进来的人的id，说明这消息是他自己发的
			uid := chatGroups[chatGroupID].Message[i].UserID
			conn.WriteJSON(messageToSend{
				UserID:   uid,
				UserName: dao.UserID2NameMap[uid],
				Message:  chatGroups[chatGroupID].Message[i].Message,
				IsSelf:   uid == userID,
			})
		}
	}

	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getChatGroupID(userID string, targetUserID string) string {
	if userID >= targetUserID {
		userID, targetUserID = targetUserID, userID
	}
	return userID + "_" + targetUserID
}

func UnRegister(userID string, targetUserID string) error {
	chatGroupID := getChatGroupID(userID, targetUserID)
	delete(chatGroups[chatGroupID].Conns, userID)
	return nil
}

// RecveMessage 接收消息，并且发送给这个会话组中的所有成员
func RecvMessage(userID string, targetUserID string, message string) error {
	chatGroupID := getChatGroupID(userID, targetUserID)
	newMessage := messageT{
		UserID:  userID,
		Message: message,
	}
	chatGroups[chatGroupID].Message = append(chatGroups[chatGroupID].Message, newMessage)
	for ID, conn := range chatGroups[chatGroupID].Conns {
		log.Println("sending from ", userID, "to", ID)
		newMessageToSend := messageToSend{
			UserID:   userID,
			UserName: dao.UserID2NameMap[userID],
			Message:  message,
			IsSelf:   ID == userID,
		}
		messageB, err := json.Marshal(newMessageToSend)
		if err != nil {
			log.Println("[SendMessage]", err)
			continue
		}
		log.Println(err, userID, targetUserID, string(messageB))
		err = conn.WriteMessage(websocket.TextMessage, messageB)
		if err != nil {
			log.Println("[SendMessage]", err)
			continue
		}
	}
	return nil
}

package chatManager

import (
	"encoding/json"
	"liansushe/dao"
	"log"

	"github.com/gorilla/websocket"
)

// 表示两个人的会话
// type chatGroup struct {
// 	Conns   map[string]*websocket.Conn
// 	Message []messageT
// }

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

// var chatGroups = make(map[string]*chatGroup)
var userID2Conn = make(map[string]*websocket.Conn)

// map[userA]userIDs 和A有聊天的所有人
var chatMates = make(map[string][]string)

// map[chatGroupID][]messageT
var messages = make(map[string][]messageT)

// Register 记录Conn
// todo: 现在注册的时候不知道是和谁讲话，只要保存下conn即可
// 接收消息时，如果没有chatGroup，则新建。需要区分是哪个chatGroup，在那个group下广播
func Register(userID string, conn *websocket.Conn) error {
	log.Println("[Register]", userID)
	userID2Conn[userID] = conn
	pushMessages(userID)
	return nil
}

func createChatGroupNx(userID string, targetUserID string) string {
	chatGroupID := getChatGroupID(userID, targetUserID)
	if _, ok := messages[chatGroupID]; ok {
		return chatGroupID
	} else {
		messages[chatGroupID] = make([]messageT, 0)
		chatMates[userID] = append(chatMates[userID], targetUserID)
		chatMates[targetUserID] = append(chatMates[targetUserID], userID)
		return chatGroupID
	}
}

func pushMessages(userID string) {
	// 对每个会话
	for _, targetUserID := range chatMates[userID] {
		chatGroupID := getChatGroupID(userID, targetUserID)
		// 推送最多5条消息
		if _, ok := messages[chatGroupID]; !ok {
			continue
		}
		length := len(messages[chatGroupID])
		for i := max(0, length-5); i < length; i++ {
			// uid是当时这条消息的发送者，userid是我们要推送给的对象
			uid := messages[chatGroupID][i].UserID
			err := userID2Conn[userID].WriteJSON(messageToSend{
				UserID:   uid,
				UserName: dao.UserID2NameMap[uid],
				Message:  messages[chatGroupID][i].Message,
				IsSelf:   uid == userID,
			})
			if err != nil {
				log.Println("[pushMessages]", err)
			}
		}
	}
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

func UnRegister(userID string) error {
	delete(userID2Conn, userID)
	return nil
}

// RecveMessage 接收消息，并且发送给这个会话组中的所有成员
func RecvMessage(userID string, targetUserID string, message string) error {
	newMessage := messageT{
		UserID:  userID,
		Message: message,
	}
	chatGroupID := createChatGroupNx(userID, targetUserID)
	messages[chatGroupID] = append(messages[chatGroupID], newMessage)
	sendMessage(userID, userID, message)
	sendMessage(userID, targetUserID, message)
	return nil
}

// 如果toer已注册进来的话，就发送一条消息给toer，这条消息是writer写的
func sendMessage(writer string, toer string, message string) error {
	newMessageToSend := messageToSend{
		UserID:   writer,
		UserName: dao.UserID2NameMap[writer],
		Message:  message,
		IsSelf:   writer == toer,
	}
	messageB, err := json.Marshal(newMessageToSend)
	if err != nil {
		log.Println("[SendMessage]", err)
		return err
	}
	log.Println(string(messageB))
	// 如果这人没注册进来就不给他发了
	if _, ok := userID2Conn[toer]; !ok {
		return nil
	}
	err = userID2Conn[toer].WriteMessage(websocket.TextMessage, messageB)
	if err != nil {
		log.Println("[SendMessage]", err)
		return err
	}
	return nil
}

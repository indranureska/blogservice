package function

import (
	serviceConst "github.com/indranureska/service/rest/common"
)

var messageMap map[string]string

// TODO: to get messages text from database
func InitServiceMessages() {
	// Initialize messages map
	messageMap = make(map[string]string)

	// Add messages to map
	messageMap[serviceConst.USER_NOT_FOUND_MSG_KEY] = serviceConst.RECORD_NOT_FOUND_MSG_DEF_TEXT
	messageMap[serviceConst.USER_NOT_FOUND_MSG_KEY] = serviceConst.USER_NOT_FOUND_MSG_DEF_TEXT
	messageMap[serviceConst.SERVICE_NOT_AVAILABLE_MSG_KEY] = serviceConst.SERVICE_NOT_AVAILABLE_MSG_DEF_TEXT
	messageMap[serviceConst.DB_CONNECT_FAILED_MSG_KEY] = serviceConst.DB_CONNECT_FAILED_MSG_DEF_TEXT
	messageMap[serviceConst.INVALID_REQUEST_PAYLOAD_MSG_KEY] = serviceConst.INVALID_REQUEST_PAYLOAD_MSG_DEF_TEXT
	messageMap[serviceConst.USER_INFO_EMPTY_MSG_KEY] = serviceConst.USER_INFO_EMPTY_MSG_DEF_TEXT
	messageMap[serviceConst.USER_EMAIL_BLANK_MSG_KEY] = serviceConst.USER_EMAIL_BLANK_MSG_DEF_TEXT
	messageMap[serviceConst.USER_PASSWORD_BLANK_MSG_KEY] = serviceConst.USER_PASSWORD_BLANK_MSG_DEF_TEXT
	messageMap[serviceConst.USER_LOGIN_FAILED_MSG_KEY] = serviceConst.USER_LOGIN_FAILED_MSG_DEF_TEXT
	messageMap[serviceConst.USER_LOGOUT_FAILED_MSG_KEY] = serviceConst.USER_LOGOUT_FAILED_MSG_DEF_TEXT
}

func ConstructServiceMessage(messageKey string) (messageText string) {
	return messageMap[messageKey]
}

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
	messageMap[serviceConst.USER_NOT_FOUND_MSG_KEY] = serviceConst.USER_NOT_FOUND_MSG_DEF_TEXT
	messageMap[serviceConst.SERVICE_NOT_AVAILABLE_MSG_KEY] = serviceConst.SERVICE_NOT_AVAILABLE_MSG_DEF_TEXT
	messageMap[serviceConst.DB_CONNECT_FAILED_MSG_KEY] = serviceConst.DB_CONNECT_FAILED_MSG_DEF_TEXT

	// const RECORD_NOT_FOUND_MSG_KEY = "record.not.found"
	// const RECORD_NOT_FOUND_MSG_DEF_TEXT = "Record not found"

	// const INVALID_RECORD_MSG_KEY = "invalid.record"
	// const INVALID_RECORD_MSG_DEF_TEXT = "Invalid record"

	// const INVALID_REQUEST_PAYLOAD_MSG_KEY = "invalid.request.payload"
	// const INVALID_REQUEST_PAYLOAD_MSG_DEF_TEXT = "Invalid request payload"

	// const USER_CREATION_FAILED_MSG_KEY = "user.creation.failed"
	// const USER_CREATION_FAILED_MSG_DEF_TEXT = "User creation failed"

	// const USER_EMAIL_ADDR_EXIST_MSG_KEY = "user.email.address.exist"
	// const USER_EMAIL_ADDR_EXIST_MSG_DEF_TEXT = "User email address exist"

	// const USER_UPDATE_FAILED_MSG_KEY = "user.update.failed"
	// const USER_UPDATE_FAILED_MSG_DEF_TEXT = "User update failed"

	// const USER_DELETE_FAILED_MSG_KEY = "user.delete.failed"
	// const USER_DELETE_FAILED_MSG_DEF_TEXT = "User delete failed"
}

// TODO: To get message text from map, set default text if it's not exist
func ConstructServiceMessage(messageKey string) (messageText string) {
	return messageMap[messageKey]
}

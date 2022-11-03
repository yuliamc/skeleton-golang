package messages

import "modalrakyat/skeleton-golang/pkg/utils/constant"

// Code represent message
const (
	MSG_CODE_NULL constant.ReserveMessageCode = iota + 20000
	MSG_CODE_INSERT_SUCCESS
	MSG_CODE_DELETE_SUCCESS
	MSG_CODE_NOTIFICATION_ACCEPTED
)

// KEYS translate error code to i18n key
var KEYS = map[constant.ReserveMessageCode]string{
	MSG_CODE_NULL:                  "MSG_CODE_NULL",
	MSG_CODE_INSERT_SUCCESS:        "MSG_CODE_INSERT_SUCCESS",
	MSG_CODE_DELETE_SUCCESS:        "MSG_CODE_DELETE_SUCCESS",
	MSG_CODE_NOTIFICATION_ACCEPTED: "MSG_CODE_NOTIFICATION_ACCEPTED",
}

func NewMessageCode(value string) int {
	for i, v := range KEYS {
		if v == value {
			return int(i)
		}
	}
	panic("Message code not found")
}

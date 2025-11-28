package common

import "sync/atomic"

const (
	ChatTypePrivate = 0
	ChatTypeGroup   = 1
	ChatTypePublic  = 2
)

func GenerateChatID(chatType int, counter uint64) int64 {
	localID := atomic.AddUint64(&counter, 1)
	return int64((uint64(chatType) << 62) | localID)
}

func GetChatType(chatID int64) int {
	return int(uint64(chatID) >> 62)
}

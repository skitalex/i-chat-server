package chatsrv

import (
	"sync"
)

type chat struct {
	m       sync.RWMutex
	chatID  int64
	clients map[string]*client
}

func newChat(chatID int64) *chat {
	return &chat{
		chatID:  chatID,
		clients: make(map[string]*client),
	}
}

func (c *chat) addClient(client *client) {
	c.m.Lock()
	defer c.m.Unlock()
	c.clients[client.id] = client
}

func (c *chat) removeClient(id string) {
	c.m.Lock()
	defer c.m.Unlock()
	delete(c.clients, id)
}

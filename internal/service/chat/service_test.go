package chatsrv

import (
	msgdomain "chatsrv/internal/domain/msg"
	"golang.org/x/net/websocket"
	"testing"
)

func Test_chatService_handleJoinChat(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ws      *websocket.Conn
		msg     msgdomain.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c chatService
			gotErr := c.handleJoinChat(tt.ws, tt.msg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("handleJoinChat() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("handleJoinChat() succeeded unexpectedly")
			}
		})
	}
}

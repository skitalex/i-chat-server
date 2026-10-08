package routes

import (
	"chatsrv/internal/controller"
	"net/http"
	"net/url"

	"golang.org/x/net/websocket"
)

func InitRoutes(ctrl controller.ChatController) *http.ServeMux {
	mux := http.NewServeMux()

	wsServer := &websocket.Server{
		Handler: func(ws *websocket.Conn) {
			ctrl.HandleWebSocket(ws)
		},
		Handshake: func(config *websocket.Config, r *http.Request) error {
			originStr := r.Header.Get("Origin")
			if originStr != "" {
				if origin, err := url.Parse(originStr); err == nil {
					config.Origin = origin
				}
			}
			return nil
		},
	}

	mux.Handle("/ws", wsServer)
	mux.HandleFunc("/api/v1/chats", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ctrl.GetChats(w, r)
		case "POST":
			ctrl.CreateChat(w, r)
		}
	})
	
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	

	return mux
}

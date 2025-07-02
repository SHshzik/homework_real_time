package websocket

import (
	"net/http"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/usecase"
	"github.com/gorilla/websocket"
)

var (
	channelBufferSize = 256
	readBufferSize    = 1024
	writeBufferSize   = 1024
	upgrader          = websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Handler struct {
	hub                 *domain.Hub
	notificationUseCase *usecase.NotificationUseCase
}

func NewHandler(hub *domain.Hub, notificationUseCase *usecase.NotificationUseCase) *Handler {
	return &Handler{
		hub:                 hub,
		notificationUseCase: notificationUseCase,
	}
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	client := domain.NewClient(h.hub, conn, make(chan []byte, channelBufferSize))
	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return nil
}

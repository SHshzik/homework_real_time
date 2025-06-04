package websocket

import (
	"log"
	"net/http"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/usecase"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
		log.Println(err)
		return err
	}
	client := domain.NewClient(h.hub, conn, make(chan []byte, 256))
	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return nil
}

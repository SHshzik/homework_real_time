package websocket

import (
	"net/http"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/usecase"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/gorilla/websocket"
)

const (
	channelBufferSize = 256
	readBufferSize    = 1024
	writeBufferSize   = 1024
)

type Handler struct {
	hub                 *domain.Hub
	notificationUseCase *usecase.NotificationUseCase
	logger              logger.Interface
}

func NewHandler(hub *domain.Hub, notificationUseCase *usecase.NotificationUseCase, logger logger.Interface) *Handler {
	return &Handler{
		hub:                 hub,
		notificationUseCase: notificationUseCase,
		logger:              logger,
	}
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	client := domain.NewClient(h.hub, conn, make(chan []byte, channelBufferSize), h.logger)
	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return nil
}

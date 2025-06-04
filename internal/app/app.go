package app

import (
	"flag"
	"log"
	"net/http"

	"github.com/SHshzik/homework_real_time/config"
	ws "github.com/SHshzik/homework_real_time/internal/controller/websocket"
	"github.com/SHshzik/homework_real_time/internal/domain"
)

var addr = flag.String("addr", ":8081", "http service address")

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	_ = cfg

	hub := domain.NewHub()
	go hub.Run()

	server := ws.NewHandler(hub, nil)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.HandleWebSocket(w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

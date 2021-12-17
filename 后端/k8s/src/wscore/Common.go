package wscore

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrader websocket.Upgrader

func init() {
	Upgrader=websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
		  return true
		},
	}
}
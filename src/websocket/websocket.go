package websocket

import(
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)  {

	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return conn, err
    }

	return conn, err
}
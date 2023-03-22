import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	// add connection to list of active connections
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}

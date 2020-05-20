package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Home")
}
func webSocketsEndpoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"Hello Sockets")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error in upgrading")
	}
	log.Println("Connecte with Client!!")
	reader(ws)
}
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("error while listioning by server")
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte(("server said Brooo") + string(p)) ); err != nil {
			log.Println("error while Writtong by server")
			return
		}
	}
}
func setupRoutes() {
	http.HandleFunc("/", homePageEndpoint)
	http.HandleFunc("/ws", webSocketsEndpoint)
}
func main() {
	setupRoutes()
	fmt.Println("Helllll")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// config file - run
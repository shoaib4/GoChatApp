package main

import (
	"fmt"
	"log"
	// "flag"
	// // "net"
	// "log"
	// "net/url"
	// "io/ioutil"
	// "net/http"
	"github.com/gorilla/websocket"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Home")
// }
// func webSocketsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w,"Hello Sockets")
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("error in upgrading")
// 	}
// 	log.Println("Connecte with Client!!")
// 	reader(ws)
// }
// func reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("error while listioning by server")
// 			return
// 		}
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println("error while Writtong by server")
// 			return
// 		}
// 	}
// }
// var addr = flag.String("addr", "localhost:8080", "http service address")
func setupRoutes() {
	// http.HandleFunc("/", homePageEndpoint)
	// http.HandleFunc("/ws", webSocketsEndpoint)
	// u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	// log.Printf("connecting to %s", u.String())
	// fmt.Println("Starting client...",u.String())
	c ,_, error := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	c.WriteMessage(websocket.TextMessage, []byte("bitchhhhh..."))
	// _, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		fmt.Println(error)
	}
	defer c.Close()
	for {
		_,mesg,err := c.ReadMessage()
		if err != nil{
			log.Println(err)
			fmt.Println(err)
			return
		}
		fmt.Println(string(mesg))

	}
	// ioutil.ReadAll(connection)
	// resp, err := http.Get("http://localhost:8080/ws")
	// defer resp.Body.Close()
	// body, err1 := ioutil.ReadAll(resp.Body)
	// if err == nil {
	// 	if err1 == nil {
	// 		fmt.Println(string(body))
	// 	} else {
	// 		fmt.Println(err1)
	// 	}
	// } else {
	// 	fmt.Println(err)
	// }
}
func main() {
	setupRoutes()
	fmt.Println("Helllll")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

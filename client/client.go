package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/gorilla/websocket"
)
func readFromServer(c *websocket.Conn){
	for {
		_, mesg, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			fmt.Println(err)
			return
		}
		fmt.Println(string(mesg))
	
	}
}
func readIOandWriteToServer(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if len(strings.Split(text, " ")) <= 1 {
			fmt.Println("Unknown command ")
			continue
		}
		c.WriteMessage(websocket.TextMessage, []byte(text))
		fmt.Println(text)
	}
}
func setupRoutes() {
	// http.HandleFunc("/", homePageEndpoint)
	// http.HandleFunc("/ws", webSocketsEndpoint)
	// u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	// log.Printf("connecting to %s", u.String())
	// fmt.Println("Starting client...",u.String())
	c, _, error := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	c.WriteMessage(websocket.TextMessage, []byte("e bitchhhhh started..."))
	// _, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		fmt.Println(error)
	}
	defer c.Close()
	go readFromServer(c)

	readIOandWriteToServer(c)
	

}
func main() {
	setupRoutes()
	fmt.Println("Helllll")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

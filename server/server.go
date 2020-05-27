package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	// https://www.gorillatoolkit.org/pkg/websocke
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var webSocketMap map[int]*websocket.Conn = make(map[int]*websocket.Conn)

var webSocketGroupMap map[int][]int = make(map[int][]int)

var nextSocketNumber int = 1

func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Home")
}
func webSocketsEndpoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w,"Hello Sockets")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error in upgrading")
		return
	}
	webSocketMap[nextSocketNumber] = ws
	nextSocketNumber++
	log.Println("Connecte with Client!!", len(webSocketMap))

	go reader(ws)
}

func getCommandAndMessage(text string) (string, string) {
	command := strings.Split(text, " ")[0]
	msg := strings.Join(strings.Split(text, " ")[1:], " ")
	return command, msg
}
func deleteWebSocket(conn *websocket.Conn) {
	for key, val := range webSocketMap {
		if val == conn {
			log.Println("****  One connection closed  ****")
			delete(webSocketMap, key)
			break
		}
	}
}
func exicuteCommand(conn *websocket.Conn, messageType int, command string, message string) error {
	commandType := int32(command[0])
	if commandType != 'c' && commandType != 'g' && commandType != 'e' {
		return errors.New("first letter is not c or g")
	}
	if commandType == 'e' {
		return nil
	}
	id, err := strconv.Atoi(string(command[1]))
	if err != nil {
		return err
	}
	if webSocketMap[id] != nil {
		if commandType == 'c' {
			if err := webSocketMap[id].WriteMessage(messageType, []byte(message)); err != nil {
				log.Println("error while Writtong by server")
				return err
			}
		}
		if commandType == 'g' {
			for _, valID := range webSocketGroupMap[id] {
				if err := webSocketMap[valID].WriteMessage(messageType, []byte(message)); err != nil {
					log.Println("error while Writtong by server")
					return err
				}
			}
		}

	} else {
		return errors.New("Socket with that Id not found")
	}
	return nil

}
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			deleteWebSocket(conn)
			return
		}
		if err := conn.WriteMessage(messageType, []byte(("server echoing Brooo :-")+string(p))); err != nil {
			log.Println("error while Writtong by server")
			return
		}
		command, message := getCommandAndMessage(string(p))
		fmt.Println("read this   :: ", command, message)

		err = exicuteCommand(conn, messageType, command, message)
		if err != nil {
			fmt.Println(err)
			conn.WriteMessage(messageType, []byte("Wrong command fount !!!!"))
		}

	}
}
func setupRoutes() {
	http.HandleFunc("/", homePageEndpoint)
	http.HandleFunc("/ws", webSocketsEndpoint)
}
func main() {
	webSocketGroupMap[2] = append(webSocketGroupMap[2], 1)
	webSocketGroupMap[2] = append(webSocketGroupMap[2], 3)
	setupRoutes()
	fmt.Println("Helllll")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// config file - run

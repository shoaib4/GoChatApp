# GoChatApp
This is my first golang project, I created it with minimalistic code and only implemented the nesseary functionality. I tryed to male it as simple as possible. 

# In Client
* Client needs to innitiate the socker connection by hitting some url (ws://localhost:8080/ws) 
```golang
c, _, error := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
```

# In Server
* Lets expose an api for socket connection
```golang
http.HandleFunc("/ws", webSocketsEndpoint)
```
* When client hits server with on that url (ws://localhost:8080/ws) server will create a socket connection between that client and server.
* We are using [gorilla/socket](https://www.gorillatoolkit.org/pkg/websocke) libreary for the socket stuff.
```golang
ws, err := upgrader.Upgrade(w, r, nil)
```
* When multiple clients are connected to server we read each of them in a seperate Goroutines.

# Things I Learned
1. Golang
2. goroutines 
3. Cammand line interface in Golang
4. Socket programming in Golang   

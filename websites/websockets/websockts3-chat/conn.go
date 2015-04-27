package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// The hub.
	h *hub
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.h.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
//create the upgrade object per gorilla/sockets doc
var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

type wsHandler struct {
	h *hub
}

//chat client connects here
func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//upgrade the connection per websocket standards
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	//create the connection object
	c := &connection{
		send: make(chan []byte, 256),
		ws: ws,
		h: wsh.h}

	//store it in the hub
	c.h.register <- c

	//unregister when the connection closes
	defer func() { c.h.unregister <- c }()

	//start the writer
	go c.writer()

	c.reader()
}

package main

type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func newHub() *hub {
	return &hub{
		broadcast:   make(chan []byte),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
	}
}

func (h *hub) run() {
	for {
		select {
		//Handle data from the register channel of hub struct
		case c := <-h.register:
			h.connections[c] = true

		//Handle data from the unregister channel of hub struct
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}

		//Handle data from the broadcast channel of hub struct
		//Receive from each client m
		case m := <-h.broadcast:
			//Loop through all the connected clients and send them the message.
			//If the client doesn't exist anymore, remove the client from the list
			for c := range h.connections {
				select {
				case c.send <- m:	//write the message into the send queue of each client and let each client writer go routine handle the channel
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}

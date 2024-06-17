package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	// socket is the web socker for this client
	socket *websocket.Conn

	// receive is a channel to receive messages from other clients
	receive chan []byte

	// room is the room this client is chatting in
	room *Room
}

func (c *Client) Read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}

		c.room.forward <- msg
	}
}

func (c *Client) Write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			fmt.Print("Error: ", err)
			return
		}
	}
}

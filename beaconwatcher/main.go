package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	var in []byte
	if err := websocket.Message.Receive(ws, &in); err != nil {
		log.Error("Could not receive")
		return
	}
	log.WithFields(log.Fields{"message": string(in)}).Info("Received beacon message")
}

func main() {
	log.Info("Hosting websocket for receiving beacon messages")

	http.Handle("/beacon", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

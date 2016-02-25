package main

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/andoco/go-test-waypoints/waypoints"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Could not upgrade http.Request to a websocket")
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				break
				log.WithFields(log.Fields{"error": err}).Info("Websocket closed")
			} else {
				log.WithFields(log.Fields{"error": err}).Error("Could not read beacon message from websocket")
				break
			}
		}

		log.WithFields(log.Fields{"message": string(message)}).Debug("Received beacon message")

		var stamp waypoints.Stamp
		json.Unmarshal(message, &stamp)

		log.WithFields(log.Fields{"waypointStamp": stamp}).Info("Received waypoint stamp")
	}

	log.Info("Finished request")
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Hosting websocket for receiving beacon messages")

	http.HandleFunc("/beacon", handler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

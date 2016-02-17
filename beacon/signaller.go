package beacon

import (
	log "github.com/Sirupsen/logrus"

	"golang.org/x/net/websocket"
)

type Signaller interface {
	Signal(waypointId string) error
}

type WebsocketSignaller struct {
	//Conn *websocket.Conn
}

func (s *WebsocketSignaller) Signal(waypointId string) error {
	log.WithFields(log.Fields{"waypoint": waypointId}).Info("Sending waypoint beacon signal")

	ws, err := dial()
	if err != nil {
		log.WithFields(log.Fields{"waypoint": waypointId, "error": err}).Error("Error dialing websocket connection")
		return err
	}

	if err := websocket.Message.Send(ws, waypointId); err != nil {
		log.WithFields(log.Fields{"waypoint": waypointId, "error": err}).Error("Error sending waypoint beacon signal over websocket")
	}

	return nil
}

var signaller Signaller

func init() {
	signaller = newWebsocketSignaller()
}

func newWebsocketSignaller() Signaller {
	return &WebsocketSignaller{}
}

func dial() (*websocket.Conn, error) {
	origin := "http://localhost/"
	url := "ws://localhost:12345/beacon"

	return websocket.Dial(url, "", origin)
}

func Signal(waypointId string) {
	signaller.Signal(waypointId)
}

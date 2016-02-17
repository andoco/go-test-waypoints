package beacon

import (
	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/websocket"
)

type Signaller interface {
	Signal(waypointId string) error
}

type WebsocketSignaller struct {
	//conn *websocket.Conn
}

func (s *WebsocketSignaller) Signal(waypointId string) error {
	log.WithFields(log.Fields{"waypoint": waypointId}).Info("Sending waypoint beacon signal")

	conn, err := dial()
	if err != nil {
		log.WithFields(log.Fields{"waypoint": waypointId, "error": err}).Error("Error dialing websocket connection")
		return err
	}

	//if err := websocket.Message.Send(ws, waypointId); err != nil {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(waypointId)); err != nil {
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
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:12345/beacon", nil)
	return c, err
}

func Signal(waypointId string) {
	signaller.Signal(waypointId)
}

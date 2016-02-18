package beacon

import (
	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/websocket"
)

type Signaller interface {
	Close()
	Signal(waypointId string) error
}

type WebsocketSignaller struct {
	conn *websocket.Conn
}

func (s *WebsocketSignaller) Close() {
	if err := s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		//if err := s.conn.Close(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Failed to close websocket")
	}
}

func (s *WebsocketSignaller) Signal(waypointId string) error {
	log.WithFields(log.Fields{"waypoint": waypointId}).Info("Sending waypoint beacon signal")

	s.ensureConnected()

	if err := s.conn.WriteMessage(websocket.TextMessage, []byte(waypointId)); err != nil {
		log.WithFields(log.Fields{"waypoint": waypointId, "error": err}).Error("Error sending waypoint beacon signal over websocket")
	}

	return nil
}

func (s *WebsocketSignaller) ensureConnected() error {
	if s.conn != nil {
		return nil
	}

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:12345/beacon", nil)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error dialing websocket connection")
		return err
	}

	s.conn = conn

	return nil
}

var signaller Signaller

func init() {
	signaller = newWebsocketSignaller()
}

func newWebsocketSignaller() Signaller {
	return &WebsocketSignaller{}
}

func Close() {
	signaller.Close()
}

func Signal(waypointId string) {
	signaller.Signal(waypointId)
}

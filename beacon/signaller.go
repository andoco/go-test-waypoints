package beacon

import (
	log "github.com/Sirupsen/logrus"
)

type Signaller interface {
	Signal(waypointId string) error
}

type WebsocketSignaller struct {
}

func (s *WebsocketSignaller) Signal(waypointId string) error {
	log.WithFields(log.Fields{"waypoint": waypointId}).Info("Sending waypoint beacon signal")
	return nil
}

var signaller Signaller

func init() {
	signaller = &WebsocketSignaller{}
}

func Signal(waypointId string) {
	signaller.Signal(waypointId)
}

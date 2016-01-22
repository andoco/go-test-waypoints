package waypoints

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	SetOutput(os.Stdout)
}

// Set the file that waypoint events will be logged to.
func SetOutput(file *os.File) {
	log.Out = file
}

// Record that the waypoint has been visited by logging a waypoint event.
func Visit(waypointId string) {
	log.WithFields(logrus.Fields{"waypoint": waypointId}).Info("Visited waypoint")
}

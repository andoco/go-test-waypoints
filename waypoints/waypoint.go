package waypoints

import (
	"errors"
	"os"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

type WaypointState struct {
	Id      string
	Visited uint64
}

var waypoints map[string]*WaypointState

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	SetOutput(os.Stdout)

	waypoints = make(map[string]*WaypointState)
}

// Set the file that waypoint events will be logged to.
func SetOutput(file *os.File) {
	log.Out = file
}

func Add(waypointId string) error {
	log.WithFields(logrus.Fields{"waypoint": waypointId}).Info("Adding waypoint")

	_, ok := waypoints[waypointId]
	if ok {
		log.WithFields(logrus.Fields{"waypoint": waypointId}).Error("Waypoint already exists")
		return errors.New("Waypoint has already been added")
	}

	waypoints[waypointId] = &WaypointState{Id: waypointId}

	log.WithFields(logrus.Fields{"waypoint": waypointId}).Info("Added waypoint")

	return nil
}

// Record that the waypoint has been visited by logging a waypoint event.
func Visit(waypointId string) error {
	log.WithFields(logrus.Fields{"waypoint": waypointId}).Info("Visiting waypoint")

	waypoint, ok := waypoints[waypointId]
	if !ok {
		log.WithFields(logrus.Fields{"waypoint": waypointId}).Error("Waypoint not found when visiting waypoint")
		return errors.New("Waypoint not found")
	}

	waypoint.Visited += 1
	log.WithFields(logrus.Fields{"waypoint": waypointId, "visited": waypoint.Visited}).Info("Visited waypoint")

	return nil
}
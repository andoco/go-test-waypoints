package waypoints

import (
	"errors"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

type WaypointState struct {
	Id      string
	Visited uint64
}

type Stamp struct {
	WaypointId    string
	CorrelationId string
	VisitedTime   time.Time
	Passed        bool
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
func Visit(waypointId string, correlationId string) error {
	stamp := Stamp{WaypointId: waypointId, CorrelationId: correlationId}

	log.WithFields(logrus.Fields{"waypointStamp": stamp}).Info("Visiting waypoint")

	waypoint, ok := waypoints[waypointId]
	if !ok {
		log.WithFields(logrus.Fields{"waypointStamp": stamp}).Error("Waypoint not found when visiting waypoint")
		return errors.New("Waypoint not found")
	}

	waypoint.Visited += 1

	return nil
}

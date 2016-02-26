package circuit

import (
	"errors"
	"fmt"

	"bitbucket.org/andoco/go-test-waypoints/waypoints"
	log "github.com/Sirupsen/logrus"
)

var circuits map[string]Observed

// Describes a planned sequence of visited waypoints.
type Circuit struct {
	Waypoints []string
}

func (c *Circuit) Evaluate(visited []waypoints.Stamp) error {
	if len(visited) != len(c.Waypoints) {
		return errors.New("wrong number of waypoints visited")
	}

	return nil
}

// An observed sequence of waypoints that were visited.
type Observed struct {
	Visited []waypoints.Stamp
}

func Record(stamp waypoints.Stamp) error {
	log.WithFields(log.Fields{"waypointStamp": stamp}).Debug("Recording waypoint stamp")

	circuit := circuits[stamp.CorrelationId]

	circuit.Visited = append(circuit.Visited, stamp)

	return nil
}

func GetObserved(correlationId string) (*Observed, error) {
	if val, ok := circuits[correlationId]; ok {
		return &val, nil
	}

	return nil, fmt.Errorf("Observed circuit %s does not exist", correlationId)
}

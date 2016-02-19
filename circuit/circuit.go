package circuit

import (
	"errors"

	"bitbucket.org/andoco/go-test-waypoints/waypoints"
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
	circuit := circuits[stamp.CorrelationId]

	circuit.Visited = append(circuit.Visited, stamp)

	return nil
}

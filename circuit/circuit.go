package circuit

import (
	"bitbucket.org/andoco/go-test-waypoints/waypoints"
)

var circuits map[string]Observed

// Describes a planned sequence of visited waypoints.
type Circuit struct {
	Waypoints []string
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

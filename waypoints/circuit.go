package waypoints

import "time"

// Describes a planned sequence of visited waypoints.
type Circuit struct {
	Waypoints []string
}

// An observed sequence of waypoints that were visited.
type ObservedCircuit struct {
	Visited []Stamp
}

type Stamp struct {
	WaypointId    string
	CorrelationId string
	VisitedTime   time.Time
	Passed        bool
}

func Record(stamp Stamp) error {
	return nil
}

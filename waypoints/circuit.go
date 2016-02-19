package waypoints

// Describes a planned sequence of visited waypoints.
type Circuit struct {
	Waypoints []string
}

// An observed sequence of waypoints that were visited.
type ObservedCircuit struct {
	Visited []Stamp
}

func Record(stamp Stamp) error {
	return nil
}

package circuit

import (
	"testing"

	"bitbucket.org/andoco/go-test-waypoints/waypoints"
)

func TestEvaluate(t *testing.T) {
	stamps := []waypoints.Stamp{
		waypoints.Stamp{WaypointId: "waypoint1"},
		waypoints.Stamp{WaypointId: "waypoint2"},
	}

	circuit := Circuit{Waypoints: []string{"waypoint1", "waypoint2"}}

	err := circuit.Evaluate(stamps)

	if err != nil {
		t.Error(err)
	}
}

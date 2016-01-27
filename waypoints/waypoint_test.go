package waypoints

import "testing"

func TestAdd(t *testing.T) {
	Add("test-waypoint")

	waypoint, ok := waypoints["test-waypoint"]

	if !ok {
		t.Error("Should add waypoint to map")
	}

	if waypoint.Id != "test-waypoint" {
		t.Error("Id should match")
	}
}

func TestAddDuplicate(t *testing.T) {
	waypointId := "test-waypoint"
	Add(waypointId)
	err := Add(waypointId)

	if err == nil {
		t.Error("Should return error")
	}

	if len(waypoints) > 1 {
		t.Error("Should not add duplicate")
	}
}

func TestVisit(t *testing.T) {
	waypointId := "test-waypoint"
	Add(waypointId)

	waypoint, ok := waypoints[waypointId]
	if !ok {
		t.Error("Should have waypoint")
	}

	Visit(waypointId)

	if waypoint.Visited != 1 {
		t.Error("Should increment Visited", waypoint.Visited)
	}
}

func TestVisitUnknownWaypoint(t *testing.T) {
	err := Visit("unknown-waypoint")

	if err == nil {
		t.Error("Should return error")
	}
}

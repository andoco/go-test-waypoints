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

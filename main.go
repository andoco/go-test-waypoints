package main

import (
	"os"

	"bitbucket.org/andoco/go-test-waypoints/waypoints"
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	waypoints.Visit("test-waypoint")
}

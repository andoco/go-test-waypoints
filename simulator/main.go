package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"bitbucket.org/andoco/go-test-waypoints/waypoints"
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	waypoints.SetOutput(os.Stdout)
}

func simulateWaypoints(interrupt <-chan os.Signal, numWaypoints int, numVisits int, done chan struct{}) {
	log.WithFields(log.Fields{"numWaypoints": numWaypoints, "numVisits": numVisits}).Info("Simulating waypoints")

	defer close(done)

	for i := 0; i < numWaypoints; i++ {
		waypoints.Add(fmt.Sprintf("test-waypoint-%d", i))
	}

	visitCount := 0
	var timeout <-chan time.Time
	timeout = time.After(1 * time.Second)

	for i := 0; i < numVisits; i++ {
		select {
		case <-interrupt:
			break
		case <-timeout:
			timeout = time.After(1 * time.Second)
			waypoints.Visit(fmt.Sprintf("test-waypoint-%d", int(math.Mod(float64(visitCount), float64(numWaypoints)))))
			if visitCount == numVisits*numWaypoints {
				break
			}
		}
	}

	log.Info("Finished simulating waypoints")
}

func main() {
	var (
		numWaypoints = kingpin.Arg("num", "The number of waypoints to simulate.").Default("1").Int()
		numVisits    = kingpin.Arg("visits", "The number of visits to make to each waypoint.").Default("1").Int()
	)

	kingpin.Parse()

	done := make(chan struct{})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go simulateWaypoints(interrupt, *numWaypoints, *numVisits, done)

	<-done

	log.Info("Exiting")
}

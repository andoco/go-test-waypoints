package main

import (
	"fmt"
	"os"
	"sync"
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

func simulateWaypoints(c <-chan struct{}, numWaypoints int, wg sync.WaitGroup) {
	log.WithFields(log.Fields{"numWaypoints": numWaypoints}).Info("Simulating waypoints")

	for i := 0; i < numWaypoints; i++ {
		waypoints.Add(fmt.Sprintf("test-waypoint-%d", i))
	}

	for i := 0; ; i++ {
		waypoints.Visit("test-waypoint-0")
		time.Sleep(1 * time.Second)
	}

	log.Info("Finished simulating waypoints")
	wg.Done()
}

func main() {
	var (
		numWaypoints = kingpin.Arg("num", "The number of waypoints to simulate.").Int()
	)

	kingpin.Parse()

	var wg sync.WaitGroup

	c := make(chan struct{})

	wg.Add(1)
	go simulateWaypoints(c, *numWaypoints, wg)

	wg.Wait()
}

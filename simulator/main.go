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

func simulateWaypoints(c <-chan struct{}, wg sync.WaitGroup) {
	log.Info("Simulating waypoints")

	for i := 0; ; i++ {
		waypoints.Visit(fmt.Sprintf("test-waypoint-%d", i))
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

	log.Info("Simulating %d waypoints", numWaypoints)

	var wg sync.WaitGroup

	c := make(chan struct{})

	wg.Add(1)
	go simulateWaypoints(c, wg)

	wg.Wait()
}

package main

import (
	"fmt"
	"os"
	"sync"
	"time"

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
	var wg sync.WaitGroup

	c := make(chan struct{})

	wg.Add(1)
	go simulateWaypoints(c, wg)

	wg.Wait()
}

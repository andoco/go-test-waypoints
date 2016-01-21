package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	log.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")
}

package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/fsnotify.v1"

	"bitbucket.org/andoco/go-test-waypoints/beacon"
)

func logModified(ev fsnotify.Event) {
	if ev.Op&fsnotify.Write == fsnotify.Write {
		log.Println("MODIFIED")
	}
}

func decodeLog() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		for k := range v {
			if k == "waypoint" {
				log.Println("Found waypoint event", v[k])
				beacon.Signal(v[k].(string))
			}
		}

		if err := enc.Encode(&v); err != nil {
			log.Fatal(err)
		}
	}
}

func watchLogfile(logfile string) {
	log.Printf("logfile = %s", logfile)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(logfile)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Events:
			log.Println("event:", ev)
			logModified(ev)
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func main() {
	var (
		logfile = kingpin.Arg("logfile", "The path to the logfile to watch.").String()
	)

	kingpin.Parse()

	if *logfile != "" {
		watchLogfile(*logfile)
	} else {
		decodeLog()
	}
}

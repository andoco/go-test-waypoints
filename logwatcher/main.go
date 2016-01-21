package main

import (
	"gopkg.in/fsnotify.v1"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

func logModified(ev fsnotify.Event) {
	if ev.Op & fsnotify.Write == fsnotify.Write {
		log.Println("MODIFIED")
	}
}

func main() {
	var (
		logfile = kingpin.Arg("logfile", "The path to the logfile to watch.").Required().String()
	)

	kingpin.Parse()

	log.Printf("logfile = %s", *logfile)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(*logfile)
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

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"golang-kafka-mongodb-watcher/Utils"

	"github.com/radovskyb/watcher"
)

func main() {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	//w.SetMaxEvents(1)

	// notify rename ,move,Create , Write events.
	w.FilterOps(watcher.Rename, watcher.Move, watcher.Create, watcher.Write)

	go func() {
		for {
			select {
			case event := <-w.Event:

				if strings.HasSuffix(event.Path, "txt") { // return true or false
					go Utils.ReadTxt(event.Path)
				} else if strings.HasSuffix(event.Path, "csv") {
					go Utils.ReadCsv(event.Path)
				} else if strings.HasSuffix(event.Path, "docx") {
					go Utils.ReadDoc(event.Path)
				} else if strings.HasSuffix(event.Path, "pdf") {
					go Utils.ReadPdf(event.Path)
				} else if strings.HasSuffix(event.Path, "xlsx") {
					go Utils.ReadXlsx(event.Path)
				}

				//fmt.Printf("The event is -> %s", event.Path) // Print the event's info.
				//fmt.Println()
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()
	// Watch this folder for changes.
	//  if err := w.Add("."); err != nil {
	//  	log.Fatalln(err)
	//  }

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive("/home/test/Desktop/LZ"); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	// Trigger 2 events after watcher started.
	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Create, nil)
		w.TriggerEvent(watcher.Remove, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

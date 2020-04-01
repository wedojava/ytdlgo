package main

import (
	"log"
	"time"

	"github.com/wedojava/ytdlgo/cmd/commons"
	"github.com/wedojava/ytdlgo/cmd/ytdlgo"
)

func main() {
	go server()
	// Just download the list one time.
	// getNow()
}

func server() {
	for {
		now := time.Now()
		// every 15' start workflow
		if now.Minute() == 15 {
			root, err := commons.PathGenAsDate()
			if err != nil {
				log.Fatal(err)
			}
			ytdlgo.DownloadConfOnce("", "", root)
		} else {
			time.Sleep(1 * time.Minute) // sleep 1 minute
		}
		commons.RemoveRoutine("")
	}
}

func getNow() {
	root, err := commons.PathGenAsDate()
	if err != nil {
		log.Fatal(err)
	}
	ytdlgo.DownloadConfOnce("", "", root)

}

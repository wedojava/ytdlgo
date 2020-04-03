package main

import (
	"log"
	"os"
	"time"

	"github.com/wedojava/ytdlgo/cmd/commons"
	"github.com/wedojava/ytdlgo/cmd/ytdlgo"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "1" {
			println("Download service start now!")
			go server()
		}
		if os.Args[1] == "2" {
			println("Download service start now!")
			getNow()
		}
	}
	if len(os.Args) == 1 {
		println("Deal with the download list once right now...")
		getNowDirectly()
	}
	os.Exit(3)
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

func getNowDirectly() {
	root, err := commons.PathGenAsDate()
	if err != nil {
		log.Fatal(err)
	}
	ytdlgo.DownloadWatchLinks("", "", root)

}

func getNow() {
	root, err := commons.PathGenAsDate()
	if err != nil {
		log.Fatal(err)
	}
	ytdlgo.DownloadConfOnce("", "", root)

}

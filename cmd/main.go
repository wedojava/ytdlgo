package main

import (
	"os"
	"time"

	"github.com/wedojava/ytdlgo/internal/ytdlgo"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "1" {
			println("Download service start now!")
			server()
		}
		if os.Args[1] == "2" {
			println("Download once start right now!")
			ytdlgo.DownloadConfOnce("", "", "Downloads")
		}
	}
	if len(os.Args) == 1 {
		println("Deal with the download list once right now...")
		ytdlgo.DownloadWatchLinks("list.txt", "", "Downloads")
	}
	os.Exit(3)
}

func server() {
	for {
		now := time.Now()
		// every 15' start workflow
		if now.Minute() == 15 {
			ytdlgo.DownloadConfOnce("", "", "Downloads")
		} else {
			time.Sleep(1 * time.Minute) // sleep 1 minute
		}
		// TODO: Need to rewrite method below
		// err := commons.RemoveRoutine("")
		// if err != nil {
		//         log.Fatal(err)
		// }
	}
}

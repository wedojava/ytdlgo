package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/wedojava/ytdlgo/internal/ytdlgo"
)

func main() {
	// Download video url list directly right now
	if len(os.Args) == 1 {
		println("[***********************************************]")
		println("[*] Deal with the download list once right now...")
		println("[***********************************************]")
		ytdlgo.DownloadWatchLinks("list.txt", "", "Downloads")
	} else { // Download channel or playlist url list for times.
		numOfTimes := os.Args[1]
		n, err := strconv.Atoi(numOfTimes)
		if err != nil {
			println("[------------------------]")
			fmt.Println("[-] ", numOfTimes, " is not an integer.")
			println("[------------------------]")
			log.Fatal(err)
		}
		if n == 0 {
			println("[******************************]")
			println("[*] Download service start now!")
			println("[******************************]")
			server()
		}
		for i := 0; i < n; i++ {
			println("[*******************************************]")
			fmt.Printf("[*] Try %d times download job start just now!\n", n)
			println("[*******************************************]")
			ytdlgo.DownloadConfOnce("", "", "Downloads")
		}
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

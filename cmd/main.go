package main

import (
	"log"
	"path/filepath"

	"github.com/wedojava/ytdlgo/cmd/ytdlgo"
)

func main() {
	links, err := ytdlgo.GetLinks(filepath.Join("../", "configs", "channelmap.txt"))
	if err != nil {
		log.Fatal(err)
	}

	go ytdlgo.Service(links)

	// watches := ytdlgo.GetWatches("https://www.youtube.com/channel/UCKz6Q5oM_SEO4oIFCiAg-jw/videos")
	// for _, v := range watches {
	//         ytdlgo.Save("https://www.youtube.com/watch?v=" + v)
	// }
}

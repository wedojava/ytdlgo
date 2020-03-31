package main

import (
	"log"
	"path/filepath"

	"github.com/wedojava/ytdlgo/cmd/ytdlgo"
)

func main() {
	// "gbk" is the default setting, so, "" is also right below, be notice, if your txt file be written in windows system, it must set the right code format as your local language set.
	links, err := ytdlgo.GetLinks(filepath.Join("../", "configs", "channelmap.txt"), "gbk")
	if err != nil {
		log.Fatal(err)
	}

	go ytdlgo.Service(links)

	// watches := ytdlgo.GetWatches("https://www.youtube.com/channel/UCKz6Q5oM_SEO4oIFCiAg-jw/videos")
	// for _, v := range watches {
	//         ytdlgo.Save("https://www.youtube.com/watch?v=" + v)
	// }
}

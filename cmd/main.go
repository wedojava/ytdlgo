package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	zlog "github.com/rs/zerolog/log"
	"github.com/rylio/ytdl"
)

func main() {
	watches := getWatches("https://www.youtube.com/channel/UCKz6Q5oM_SEO4oIFCiAg-jw/videos")
	for _, v := range watches {
		save("https://www.youtube.com/watch?v=" + v)
	}
}

func getWatches(ytchannel string) []string {
	var vs []string
	// Request the HTML page.
	res, err := http.Get(ytchannel)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var re = regexp.MustCompile(`(?m)(?P<prefix>/watch\?v=)(?P<id>[a-zA-Z0-9_-]*)`)
	for _, match := range re.FindAllStringSubmatch(string(body), -1) {
		if len(match[2]) != 11 {
			continue
		} else {
			vs = append(vs, match[2])
		}
	}
	vs = removeDuplicateElement(vs)
	// printSlice(vs)

	return vs
}

// func printSlice(x []string) {
//         fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
// }

func removeDuplicateElement(items []string) []string {
	result := make([]string, 0, len(items))
	temp := map[string]struct{}{}
	for _, item := range items {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// save will download videos from youtube, and save details at the same time.
func save(url string) {
	client := ytdl.Client{
		HTTPClient: http.DefaultClient,
		Logger:     zlog.Logger,
	}
	info, err := client.GetVideoInfo(url)
	// info judgement
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	if !just1day(info.DatePublished) {
		return
	}
	// path prepare
	root, err := pathGen()
	if err != nil {
		log.Fatal(err)
	}
	title := info.Title
	pv := filepath.Join(root, title+".mp4")
	pt := filepath.Join(root, title+".txt")
	// save
	vfile, _ := os.Create(pv)
	tfile, _ := os.Create(pt)
	defer vfile.Close()
	defer tfile.Close()

	_, err = tfile.WriteString(info.Description)
	if err != nil {
		log.Fatal(err)
	}

	// 1 -> 22 -> 720p, 3 -> 37 -> 1080p
	// Look up the number and itag at itag.go and format_list_test.go
	// client.Download(info, info.Formats.Best(ytdl.FormatResolutionKey)[0], file)
	err = client.Download(info, info.Formats[1], vfile)
	if err != nil {
		log.Fatal(err)
	}
}

func just1day(t time.Time) bool {
	return time.Now().Unix()-t.Unix() < 24*60*60
}

func pathGen() (s string, err error) {
	b := time.Now()
	s1 := fmt.Sprintf("%d", b.Year())
	s2 := fmt.Sprintf("%02d%02d", b.Month(), b.Day())
	s = filepath.Join(s1, s2)
	if !Exists(s) {
		err = os.MkdirAll(s, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	return
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

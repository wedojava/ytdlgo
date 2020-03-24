package ytdlgo

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/rylio/ytdl"
	"github.com/wedojava/ytdlgo/cmd/commons"
)

func Service(links []string) {
	for {
		for _, link := range links {
			// printf error occur while debug, weird problom
			watches := GetWatches(link)
			for _, watch := range watches {
				Save("https://wwww.youtube.com/watch?v=" + watch)
			}
		}
	}
}

func GetLinks(filename string) (ls []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		ls = append(ls, GetUrl(string(a))[1])
		// fmt.Println(string(a))
	}
	return
}

// GetUrl will get the url from str, return as ["sth", "https://...."]
// profix string and url must split by ` `, `|`, `:`, `：`
func GetUrl(str string) (rt []string) {
	var re = regexp.MustCompile(`(?m)(?P<tag>.*)[\||\s|:|：](?P<url>https://.*\b)`)
	a := re.FindStringSubmatch(str)
	rt = append(rt, a[1], a[2])
	return
}

func GetWatches(ytchannel string) []string {
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
	vs = commons.StrSliceDeDupl(vs)
	// printSlice(vs)

	return vs
}

// save will download videos from youtube, and save details at the same time.
func Save(url string) {
	// client := ytdl.Client{
	//         HTTPClient: http.DefaultClient,
	//         Logger:     zlog.Logger,
	// }
	// info, err := client.GetVideoInfo(url)
	// it is dev version fit above

	// 0.6.2 version
	info, err := ytdl.GetVideoInfo(url)
	if err != nil {
		log.Fatal(err)
	}

	// info judgement
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	if !commons.Just1Day(info.DatePublished) {
		return
	}
	// path prepare
	root, err := commons.PathGenAsDate()
	if err != nil {
		log.Fatal(err)
	}
	title := info.Title
	user := info.Artist
	pv := filepath.Join(root, user, title+".mp4")
	pt := filepath.Join(root, user, title+".txt")
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
	// err = client.Download(info, info.Formats[1], vfile)  // this is dev version fit.
	err = info.Download(info.Formats[1], vfile)
	if err != nil {
		log.Fatal(err)
	}
}

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
	"unicode/utf8"

	"github.com/rylio/ytdl"
	"github.com/wedojava/ytdlgo/cmd/commons"
)

type Links struct {
	TAG string
	URL string
}

// Ytdlgo will download videos from youtube, and save details at the same time.
// @url must be youtube link of definite video. `watch?v=` at the end of url commonly.
// @tag is the subfolder name to save the downloaded files.
func Ytdlgo(url, tag, root string) {
	// 0.6.2 version
	info, err := ytdl.GetVideoInfo(url)
	// info judgement
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	//if !commons.Just1Day(info.DatePublished) {
	//	return
	//}
	// path prepare
	title := info.Title
	reg := regexp.MustCompile(`[\\/\s:*?"<>|!]`)
	rep := "${1}"
	title = reg.ReplaceAllString(title, rep)
	// user := info.Artist
	i := info.DatePublished
	vtime := fmt.Sprintf("%02d%02d", i.Month(), i.Day())
	root = filepath.Join(root, tag)
	if !commons.Exists(root) {
		os.Mkdir(root,0755)
	}
	pv := filepath.Join(root, vtime+title+".mp4")
	pt := filepath.Join(root, vtime+title+".txt")
	// save while file not exist
	if commons.Exists(pv) {
		fmt.Print("file exist and return.")
		return
	}

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

// DownloadConfOnce will run with config at `../configs/channelmap.txt` by default setting of "" for param conf, also, format is "gbk" while value is ""
func DownloadWatchLinks(conf, format, root string) {
	// "gbk" is the default setting, so, "" is also right below, be notice, if your txt file be written in windows system, it must set the right code format as your local language set.
	if conf == "" {
		conf = filepath.Join("configs", "channelmap.txt")
	}
	if format == "" {
		format = "gbk"
	}
	links, err := GetLinks(conf, format)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		Ytdlgo(link.URL, link.TAG, root)
	}
}

// DownloadConfOnce will run with config at `../configs/channelmap.txt` by default setting of "" for param conf, also, format is "gbk" while value is ""
func DownloadConfOnce(conf, format, root string) {
	// "gbk" is the default setting, so, "" is also right below, be notice, if your txt file be written in windows system, it must set the right code format as your local language set.
	if conf == "" {
		conf = filepath.Join("configs", "channelmap.txt")
	}
	if format == "" {
		format = "gbk"
	}
	links, err := GetLinks(conf, format)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		// printf error occur while debug, weird problom
		watches := GetWatches(link.URL)
		for _, watch := range watches {
			Ytdlgo("https://wwww.youtube.com/watch?v="+watch, link.TAG, root)
		}
	}
}

// GetLinks get []Links from file, param fileCode default is "gbk"
// fileCode can be set "" to omit, then it will deal file format with code gbk
func GetLinks(filename string, fileCode string) (ls []Links, err error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// fdetect := commons.FileCodeDetector(filename)

	if fileCode == "" {
		fileCode = "gbk"
	}
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(a) == 0 {
			continue
		}
		s := string(a)
		if !utf8.ValidString(s) {
			s = commons.ConvertToUtf8(s, fileCode, "utf-8")
		}
		ls = append(ls, GetUrl(s))
	}
	return
}

// GetUrl will get the url from str, return as ["sth", "https://...."]
// profix string and url must split by `|`, `:`, `：`
func GetUrl(str string) (rt Links) {
	var re = regexp.MustCompile(`(?m)\s*(?P<tag>\S*)\s*[\||:|：]\s*(?P<url>https://.*)\s*`)
	a := re.FindStringSubmatch(str)
	rt.TAG = a[1]
	rt.URL = a[2]
	return
}

// GetWatches can get []string match `/watch?v=` from body get via ytchannel.
func GetWatches(ytchannel string) []string {
	var watch_v []string
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

	var re = regexp.MustCompile(`(?m)(?P<prefix>/watch\?v=)(?P<id>.*)\&+.*`)
	for _, match := range re.FindAllStringSubmatch(string(body), -1) {
		if len(match[2]) != 11 {
			continue
		} else {
			watch_v = append(watch_v, match[2])
		}
	}
	watch_v = commons.StrSliceDeDupl(watch_v)
	// printSlice(watch_v)

	return watch_v
}

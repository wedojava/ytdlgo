package commons

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	confDir string = "configs/channelmap.txt"
)

func PathGenAsDate() (s string, err error) {
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

func FileCodeDetector(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	fdetect := StrDetector(string(fd))

	return fdetect
}

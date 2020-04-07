package commons

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
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

func RemoveFolder(folderPath string) error {
	if !Exists(folderPath) {
		return nil
	}
	err := os.RemoveAll(folderPath)
	if err != nil {
		return err
	}
	return nil
}

func RemoveRoutine(root string) error {
	a := time.Now().AddDate(0, 0, -2)
	b := fmt.Sprintf("%02d%02d", a.Month(), a.Day())
	err := RemoveFolder(filepath.Join(root, b))
	if err != nil {
		return err
	}

	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	//if err != nil && os.IsNotExist(err) {
	//	return false
	//}
	//return true
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

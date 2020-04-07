package commons

import (
	"fmt"
	"log"

	"github.com/axgle/mahonia"
	chardet2 "github.com/chennqqi/chardet"
	"github.com/gogs/chardet"
)

func PrintSlice(x []string) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// StrSliceDeDupl is used for string slice deduplication
func StrSliceDeDupl(items []string) []string {
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

func ConvertToUtf8(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func StrDetector2(s string) string {
	// detectors := chardet2.Possible([]byte(s))
	detector := chardet2.Mostlike([]byte(s))
	return detector
}

func StrDetector(s string) string {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf(
	//         "Detected charset is %s, language is %s",
	//         result.Charset,
	//         result.Language,
	// )

	return result.Charset
}

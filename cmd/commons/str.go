package commons

import "fmt"

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

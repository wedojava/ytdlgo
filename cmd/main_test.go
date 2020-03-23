package main

import (
	"reflect"
	"testing"
)

var tcs = []struct {
	ytchannel string
	want      []string
}{
	{
		"https://www.youtube.com/playlist?list=PLC0IROtmxiB_BCatMFa8_XdFTHKrIkCHY&pbjreload=10",
		[]string{"YWU4OtFpHKI"},
	},
}

func TestGeturls(t *testing.T) {
	actual := getWatches(tcs[0].ytchannel)[0]
	want := tcs[0].want[0]
	if actual != want {
		t.Errorf("Actual: %v, Excepted: %v", actual, want)
	}
}

func TestGetLinks(t *testing.T) {
	filename := "../test/test.txt"
	checkGetLinks := func(t *testing.T, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("test get txt split by ln", func(t *testing.T) {
		got, _ := getLinks(filename)
		want := []string{
			"https://tour.golang.org/moretypes/15",
			"https://www.youtube.com/watch?v=6SkbNlWMG5w",
			"https://www.youtube.com/watch?v=yT0pxoEmhGo",
		}
		checkGetLinks(t, got, want)
	})

}

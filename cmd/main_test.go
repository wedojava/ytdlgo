package main

import "testing"

var tcs = []struct {
	ytchannel string
	want      []string
}{
	{
		"https://www.youtube.com/playlist?list=PLC0IROtmxiB_BCatMFa8_XdFTHKrIkCHY&pbjreload=10",
		[]string{"https://www.youtube.com/watch?v=YWU4OtFpHKI"},
	},
}

func TestGeturls(t *testing.T) {
	actual := getWatches(tcs[0].ytchannel)[0]
	want := tcs[0].want[0]
	if actual != want {
		t.Errorf("Actual: %v, Excepted: %v", actual, want)
	}
}

package ytdlgo

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

func TestGetUrl(t *testing.T) {
	checkGetUrl := func(t *testing.T, _got, want string) {
		t.Helper()
		if _got != want {
			t.Errorf("got %v\nwant %v", _got, want)
		}
	}
	t.Run("test url getting", func(t *testing.T) {
		got := GetUrl("剪辑：https://www.youtube.com/playlist?list=PLbmJmqERD0RLGX7tdpZNlSI1cHVIHlWLT")
		want := "https://www.youtube.com/playlist?list=PLbmJmqERD0RLGX7tdpZNlSI1cHVIHlWLT"
		checkGetUrl(t, got[1], want)

	})
}

func TestGetWatches(t *testing.T) {
	actual := GetWatches(tcs[0].ytchannel)[0]
	want := tcs[0].want[0]
	if actual != want {
		t.Errorf("Actual: %v, Excepted: %v", actual, want)
	}
}

func TestGetLinks(t *testing.T) {
	filename := "../../test/test.txt"
	configfile := "../../configs/channelmap.txt"
	checkGetLinks := func(t *testing.T, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}
	t.Run("test get txt split by ln", func(t *testing.T) {
		got, _ := GetLinks(filename)
		want := []string{
			"https://tour.golang.org/moretypes/15",
			"https://www.youtube.com/watch?v=6SkbNlWMG5w",
			"https://www.youtube.com/watch?v=yT0pxoEmhGo",
		}
		checkGetLinks(t, got, want)
	})
	t.Run("test get txt in configs/channelmap.txt", func(t *testing.T) {
		got, _ := GetLinks(configfile)
		want := []string{
			"https://www.youtube.com/playlist?list=PLbmJmqERD0RI2ehh9rQfVwkm6GaksNjuJ",
			"https://www.youtube.com/playlist?list=PLbmJmqERD0RKcqyQT1DHFnVmsFG6xmjN-",
			"https://www.youtube.com/playlist?list=PLbmJmqERD0RLGX7tdpZNlSI1cHVIHlWLT",
			"https://www.youtube.com/playlist?list=PLbmJmqERD0RIt31ONx_Iveu4j2g3pIixw",
		}
		checkGetLinks(t, got, want)
	})

}

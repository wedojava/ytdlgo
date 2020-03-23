package commons

import "time"

func Just1Day(t time.Time) bool {
	return time.Now().Unix()-t.Unix() < 24*60*60
}

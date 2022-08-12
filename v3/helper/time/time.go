package time

import "time"

func NowUtc() time.Time {
	var now = time.Now().UTC()
	return now
}

func NowUtcUnix() int64 {
	return NowUtc().Unix()
}

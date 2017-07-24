package util

import "time"

func Now() int64 {
	now := time.Now().UnixNano()/10e6
	return now
}

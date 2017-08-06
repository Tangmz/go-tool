package util

import "time"

// Now返回一个13位的时间戳,即单位是毫秒
func Now() int64 {
	now := time.Now().UnixNano()/1e6
	return now
}

// Now10 返回一个10位的时间戳,即单位是秒
func Now10() int64 {
	return Now() / 1000
}

// Now13返回一个13位的时间戳,即单位是毫秒,同Now()
func Now13() int64 {
	return Now()
}
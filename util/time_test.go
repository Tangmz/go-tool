package util

import (
	"testing"
	"fmt"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(Now())
	fmt.Println(Now10())
	fmt.Println(Now13())
}

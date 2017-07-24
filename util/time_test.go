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
}

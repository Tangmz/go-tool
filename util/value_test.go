package util

import (
	"testing"
	"fmt"
	"time"
)

func TestTrimValue(t *testing.T) {
	// test string
	var strs = []string{}
	strs = TrimAryStringRepeat(strs, "", true)
	if len(strs) > 0 {
		t.Error(strs)
		return
	}

	strs = []string{"a", "b", "a"}
	strs = TrimAryStringRepeat(strs, "a", false)
	if len(strs) != 1 {
		t.Error(strs)
		return
	}
	if strs[0] != "b" {
		t.Error(strs)
		return
	}

	strs = []string{"a", "b", "a"}
	strs = TrimAryStringRepeat(strs, "", true)
	if len(strs) != 3 {
		t.Error(strs)
		return
	}
	if strs[0] != "a" || strs[1] != "b" || strs[2] != "a" {
		t.Error(strs)
		return
	}

	strs = []string{"ab", "bc", "abc", "de"}
	strs = TrimAryStringRepeat(strs, "", true)
	if len(strs) != 4 {
		t.Error(strs)
		return
	}
	if strs[0] != "ab" || strs[1] != "bc" || strs[2] != "abc" || strs[3] != "de" {
		t.Error(strs)
		return
	}

	strs = TrimAryStringRepeat(strs, "a", true)
	if len(strs) != 4 {
		t.Error(strs)
		return
	}
	if strs[0] != "b" || strs[1] != "bc" || strs[2] != "bc" || strs[3] != "de" {
		t.Error(strs)
		return
	}

	strs = []string{"b", "bc", "bc", "de"}
	strs = TrimAryStringRepeat(strs, "a", false)
	if len(strs) != 3 {
		t.Error(strs)
		return
	}
	if strs[0] != "b" || strs[1] != "bc" || strs[2] != "de" {
		t.Error(strs)
		return
	}


	// test int
	var ints = []int{1, 2}
	ints = TrimAryInt(ints, 1)
	if len(ints) != 2 {
		t.Error(ints)
		return
	}
	if ints[0] != 1 || ints[1] != 2 {
		t.Error(ints)
		return
	}

	ints = []int{1, 1, 2, 3}
	ints = TrimAryInt(ints, -1)
	if len(ints) != 3 {
		t.Error(ints)
		return
	}
	if ints[0] != 1 || ints[1] != 2 || ints[2] != 3 {
		t.Error(ints)
		return
	}


	// test int
	ints = []int{1, 1, 2, 3}
	ints = TrimAryIntRepeat(ints)
	if len(ints) != 3 {
		t.Error(ints)
		return
	}
	if ints[0] != 1 || ints[1] != 2 || ints[2] != 3 {
		t.Error(ints)
		return
	}
}

func TestUuid(t *testing.T) {
	var times int = 10
	for i := 0; i < times; i ++ {
		testUuid(t, times)
	}
}

func testUuid(t *testing.T, times int) {
	fmt.Println("times --- ", times, " START")
	uuid := UUID()
	if "" == uuid {
		t.Error(uuid)
		return
	}
	t.Log(uuid)

	// 连续生成uuid 10000个,查看是否有重复的id生成
	var length = 10000
	strs := make([]string, 0, length+10)
	for i := 0; i < length; i ++ {
		strs = append(strs, UUID())
	}
	strs = TrimAryStringRepeat(strs, "", false)
	if len(strs) != length {
		t.Error(len(strs))
	}

	// 并发生成uuid 10000个,查看是否有重复的uuid生成
	strs = []string{}
	strs = make([]string, 0, length+100)
	for i := 0; i < length; i ++ {
		go func() {
			strs = append(strs, UUID())
		}()
	}
	time.Sleep(time.Second*1)
	originLength := len(strs)
	strs = TrimAryStringRepeat(strs, "", false)
	if len(strs) != originLength {
		t.Error(len(strs))
	}
	fmt.Println("times --- ", times, " END")
}

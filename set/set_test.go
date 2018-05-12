package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	var strs = []string{}
	strs = AddToStrings(strs, "a")
	strs = AddToStrings(strs, "b")
	strs = AddToStrings(strs, "c")
	strs = AddToStrings(strs, "d")
	strs = AddToStrings(strs, "e")

	// check
	var checkStrs = []string{"a", "b", "c", "d", "e"}
	if Len(StringSlice(strs)) != len(checkStrs) {
		t.Error(len(StringSlice(strs)))
		return
	}
	if len(strs) != len(checkStrs) {
		t.Error(len(strs), strs)
		return
	}
	for i, s := range strs {
		if s != checkStrs[i] {
			t.Error(i, s)
			return
		}
	}
	fmt.Println("---------- 1 ----------")

	strs = AddToStrings(strs, "a")
	strs = AddToStrings(strs, "b")
	strs = AddToStrings(strs, "c")
	strs = AddToStrings(strs, "d")
	strs = AddToStrings(strs, "e")
	strs = AddToStrings(strs, "f")

	// check
	checkStrs = []string{"a", "b", "c", "d", "e", "f"}
	if Len(StringSlice(strs)) != len(checkStrs) {
		t.Error(len(StringSlice(strs)))
		return
	}
	if len(strs) != len(checkStrs) {
		t.Error(len(strs), strs)
		return
	}
	for i, s := range strs {
		if s != checkStrs[i] {
			t.Error(i, s)
			return
		}
	}

	fmt.Println("---------- 2 ----------")

	AddToStrings(strs, "z")
	AddToStrings(strs, "x")
	AddToStrings(strs, "c")
	AddToStrings(strs, "v")
	AddToStrings(strs, "b")

	// check
	checkStrs = []string{"a", "b", "c", "d", "e", "f"}
	if Len(StringSlice(strs)) != len(checkStrs) {
		t.Error(len(StringSlice(strs)))
		return
	}
	if len(strs) != len(checkStrs) {
		t.Error(len(strs), strs)
		return
	}
	for i, s := range strs {
		if s != checkStrs[i] {
			t.Error(i, s)
			return
		}
	}

	fmt.Println("---------- 3 ----------")

	// remove element
	strs = RemoveFromStrings(strs, "a")
	strs = RemoveFromStrings(strs, "b")
	strs = RemoveFromStrings(strs, "c")
	strs = RemoveFromStrings(strs, "d")

	// check
	checkStrs = []string{"e", "f"}
	if Len(StringSlice(strs)) != len(checkStrs) {
		t.Error(len(StringSlice(strs)))
		return
	}
	if len(strs) != len(checkStrs) {
		t.Error(len(strs), strs)
		return
	}
	for i, s := range strs {
		if s != checkStrs[i] {
			t.Error(i, s)
			return
		}
	}

	fmt.Println("---------- 4 ----------")
	RemoveFromStrings(strs, "a")
	RemoveFromStrings(strs, "b")
	RemoveFromStrings(strs, "c")
	RemoveFromStrings(strs, "d")
	RemoveFromStrings(strs, "d")
	RemoveFromStrings(strs, "e")
	RemoveFromStrings(strs, "f")

	// check
	checkStrs = []string{"e", "f"}
	if Len(StringSlice(strs)) != len(checkStrs) {
		t.Error(len(StringSlice(strs)))
		return
	}
	if len(strs) != len(checkStrs) {
		t.Error(len(strs), strs)
		return
	}
	for i, s := range strs {
		if s != checkStrs[i] {
			t.Error(i, s)
			return
		}
	}

	fmt.Println("TestSet test success!!!")
}

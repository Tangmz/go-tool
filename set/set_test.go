package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	var strs = []string{}
	strs =AddToStrings(strs, "a")
	AddToStrings(strs, "b")
	AddToStrings(strs, "c")
	AddToStrings(strs, "d")
	AddToStrings(strs, "e")

	var checkStrs = []string{"a", "b", "c", "d", "e"}
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

package util

import "testing"

func TestError(t *testing.T) {
	var err error
	if nil != err {
		t.Error(err)
		return
	}

	err = Error("%v", "halo")
	if err == nil {
		t.Error(err)
		return
	}

	if err.Error() != "halo" {
		t.Error(err)
		return
	}
}

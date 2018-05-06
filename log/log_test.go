package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func slog() {
	if true {
		LogD_(3, "slog %v", "Debug")
	}
}

func TestLog(t *testing.T) {
	D("%v", "Debug")
	I("%v", "Info")
	W("%v", "Warning")
	E("%v", "NewError")

	logger.D("logger %v", "Debug")
	logger.I("logger %v", "Info")
	logger.W("logger %v", "Warning")
	logger.E("logger %v", "Error")

	LogD_(-2, "-2 %v", "Debug")
	LogD_(-1, "-1 %v", "Debug")
	LogD_(0, "0 %v", "Debug")
	LogD_(1, "1 %v", "Debug")
	LogD_(2, "2 %v", "Debug")

	slog()

	fmt.Println("测试日志完毕")
}

func TestRedirect(t *testing.T) {
	var err error
	var filename string = "test/log/a.log"
	err = RedirectFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	info, err := os.Stat(filename)
	if err != nil {
		t.Error(err)
		return
	}
	if info.Name() != "a.log" {
		t.Error(info)
		return
	}

	D("%s", "halo")
	bys, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	var fileContent string
	fileContent = string(bys)
	if !strings.Contains(fileContent, "halo") {
		t.Error(fileContent)
		return
	}

	// redirect to file again
	err = RedirectFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	D("%s", "abc")
	bys, err = ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	fileContent = string(bys)
	if !strings.Contains(fileContent, "halo") {
		t.Error(fileContent)
		return
	}
	if !strings.Contains(fileContent, "abc") {
		t.Error(fileContent)
		return
	}
}

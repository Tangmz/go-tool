package util

import (
	"testing"
	"net/http"
	"fmt"
)

func HandleTestFunc(w http.ResponseWriter, r *http.Request) {
	msg := Map{
		"msg": "message",
	}
	fmt.Println("HandleTestFunc come in")
	w.Write([]byte(S2Json(msg)))
}

func HandleTestFuncString(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("string"))
	fmt.Println("HandleTestFuncString come in")
}

func TestRequest(t *testing.T) {
	fmt.Println("TestRequest测试开始")
	// 启动server
	go func() {
		http.HandleFunc("/message", HandleTestFunc)
		http.HandleFunc("/string", HandleTestFuncString)
		http.ListenAndServe(":8900", nil)
	}()

	var resString string
	var resMap Map
	var checkString string
	var checkMap Map
	var err error
	// 测试HTTPGetString
	resString, err = HTTPGetString("%v", "http://127.0.0.1:8900/message")
	if err != nil {
		t.Error(err)
		return
	}
	checkMap = Map{"msg": "message"}
	checkString = S2Json(checkMap)
	if resString != checkString {
		t.Error(resString)
		return
	}

	// 测试HTTPGetMap
	resMap, err = HTTPGetMap("%v", "http://127.0.0.1:8900/message")
	if err != nil {
		t.Error(err)
		return
	}
	checkMap = Map{"msg": "message"}
	if checkMap["msg"] != resMap["msg"] {
		t.Error(checkMap["msg"])
		return
	}

	// 测试HTTPPostString
	resString, err = HTTPPostString("http://127.0.0.1:8900/message", "text/html", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if resString != checkString {
		t.Error(resString)
		return
	}

	// 测试HTTPPostMap
	resMap, err = HTTPPostMap("http://127.0.0.1:8900/message", "application/json", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if resMap["msg"] != checkMap["msg"] {
		t.Error(resMap["msg"])
		return
	}

	// request for a string

	resString, err = HTTPGetString("%v", "http://127.0.0.1:8900/string")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "string" {
		t.Error(err)
		return
	}

	resMap, err = HTTPGetMap("%v", "http://127.0.0.1:8900/string")
	if err == nil {
		t.Error(err)
		return
	}

	resString, err = HTTPPostString("http://127.0.0.1:8900/string", "application/json", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "string" {
		t.Error(err)
		return
	}

	resMap, err = HTTPPostMap("http://127.0.0.1:8900/string", "application/json", nil)
	if err == nil {
		t.Error(err)
		return
	}

	fmt.Println("TestRequest测试完毕")
}

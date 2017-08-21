package http

import (
	"testing"
	"net/http"
	"github.com/tangs-drm/go-tool/util"
	"fmt"
	"time"
	"strings"
)

func HandleHalo(w http.ResponseWriter, r *http.Request) {
	msg := util.Map{"code": 0, "msg": "halo"}
	w.Write([]byte(util.S2Json(msg)))
}

func HandleString(w http.ResponseWriter, r *http.Request) {
	var err error
	_, err = w.Write([]byte("string"))
	fmt.Println("HandleString come in", err)
	return
}

// TestHttp 测试http请求是否正常
func TestHttp(t *testing.T) {
	fmt.Println("TestHttp测试 Start")
	mux := NewServerMux()
	go func() {

		// 定义路由
		mux.HandleFunc("/halo", HandleHalo)
		mux.HandleFunc("/string", HandleString)
		http.ListenAndServe(":8901", mux)
	}()

	var resMap util.Map
	var resString string

	var err error

	//resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/halo")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(resString)
	//return

	// 请求测试
	resMap, err = util.HTTPGetMap("%v", "http://127.0.0.1:8901/halo")
	if err != nil {
		t.Error(err)
		return
	}
	if resMap.Int("code") != 0 {
		t.Error(err)
		return
	}
	if resMap.String("msg") != "halo" {
		t.Error(err)
		return
	}

	resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/string")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "string" {
		t.Error(resString)
		return
	}

	// 添加过滤器
	var argFilter FilterFunc = func(w http.ResponseWriter, r *http.Request) int {
		var filter string
		filter = r.FormValue("filter")
		fmt.Printf("filter -- %v|\n", filter)
		if filter != "halo" {
			w.Write([]byte("invalid"))
			return 1
		}
		return 0
	}
	mux.FilterFunc("/usr/api",  argFilter)

	// 添加过滤器,访问其他路由
	resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/string")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "string" {
		t.Error(resString)
		return
	}

	// 添加路由,访问不受过滤
	mux.HandleFunc("/usr/ap/halo", HandleString)
	resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/usr/ap/halo")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "string" {
		t.Error(resString)
		return
	}

	// 添加路由,受过滤器影响,返回invalid
	mux.HandleFunc("/usr/api/halo", HandleString)
	resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/usr/api/halo")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "invalid" {
		t.Error(resString)
		return
	}

	//添加路由,通过过滤器
	mux.HandleFunc("/usr/api/halo", HandleString)
	resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/usr/api/halo?filter=halo")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "string" {
		t.Error(resString)
		return
	}

	//添加路由,不通过过滤器
	mux.HandleFunc("/usr/api/halo", HandleString)
	resString, err = util.HTTPGetString("%v", "http://127.0.0.1:8901/usr/api/halo?filter=halo+++")
	if err != nil {
		t.Error(err)
		return
	}
	if resString != "invalid" {
		t.Error(resString)
		return
	}

	// 添加路由, 通过路由获取map
	var pubFilter FilterFunc = func(w http.ResponseWriter, r *http.Request) int {
		if r.URL.Path != "/pub/api/map" {
			return REQUEST_RETURN
		}
		return REQUEST_CONTINUE
	}
	mux.FilterFunc("/pub/", pubFilter)
	mux.HandleFunc("/pub/api/map", HandleHalo)
	resMap, err = util.HTTPGetMap("%v", "http://127.0.0.1:8901/halo")
	if err != nil {
		t.Error(err)
		return
	}
	if resMap.Int("code") != 0 {
		t.Error(resMap.Int("code"))
		return
	}
	if resMap.String("msg") != "halo" {
		t.Error(resMap.String("msg"))
		return
	}

	fmt.Println("TestHttp测试 Start")
}

func TestFileServer(t *testing.T) {
	fmt.Println("TestFileServer 测试开始")
	var err error
	HandleFunc("/halo", HandleHalo)
	Handle("/", http.FileServer(http.Dir(".")))
	go func() {
		err = ListenAndServe(":8901", nil)
		if err != nil {
			panic(err)
		}
	}()


	time.Sleep(2*time.Second)

	var resString string

	// 检测访问文件,发送404
	resString, err = util.HTTPGetString("%v/data/notfound.txt", "http://127.0.0.1:8901")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(resString)

	if !strings.Contains(resString, "not found") {
		t.Error(resString)
		return
	}

	// 检测访问文件,成功
	resString, err = util.HTTPGetString("%v/data/test.txt", "http://127.0.0.1:8901")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(resString)
	if strings.Contains(resString, "not found") {
		t.Error(resString)
		return
	}
	fmt.Println("TestFileServer 测试完毕")
}

### http 封装

#### 说明
* 该package封装了一个简单的http请求包

* 该package保留了golang原本的http的使用风格,并在原本的代码风格上进行扩展

#### 不定时更新

1. 增加http请求路由的正则表达式

2. 增加http请求的参数方便获取,类似于Map操作一样,可以使用形如Map.String()

#### 用法,用例如下
```golang
package main

import (
	thttp "github.com/tangs-drm/go-tool/http"
	"net/http"
)

func HaloHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("halo world"))
}

func main() {
	mux := thttp.NewServerMux()

	// 普通的http路由
	mux.HandleFunc("/halo", HaloHandleFunc)

	// 添加过滤器
	var filter thttp.FilterFunc = func(w http.ResponseWriter, r *http.Request) int {
		path := r.URL.Path
		if path != "/usr/halo" {
			w.Write([]byte("wrong path"))
			return thttp.REQUEST_RETURN
		}
		return thttp.REQUEST_CONTINUE
	}
	mux.FilterFunc("/usr/halo", filter)
	mux.HandleFunc("/usr/halo", HaloHandleFunc)

	http.ListenAndServe(":8900", mux)

}

```

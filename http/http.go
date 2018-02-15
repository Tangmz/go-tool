package http

import (
	"net/http"
	"strings"
)

const (
	// 过滤器错误码
	FILTER_ERROR = 1000
)

const (
	REQUEST_CONTINUE int = 0 // 过滤器操作后继续执行
	REQUEST_RETURN   int = 1 // 过滤器操作后停机执行
)

type Mux struct {
	Handles map[string]http.Handler // 请求路由
	Filters map[string]FilterFunc   // 过滤器

	Request        *Request
	ResponseWriter ResponseWriter
}

// ServeHTTP 实现了http的ServeHTTP接口,以实现http的封装
func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// 获取过滤器,如果有则执行
	filter := mux.deliverFilter(path)
	var filterCode int
	if filter != nil {
		filterCode = filter(w, r)
		if REQUEST_RETURN == filterCode {
			return
		}
	}

	// 根据请求的路由获取执行的handler
	handler := mux.deliverHandler(path)
	if handler == nil {
		http.NotFound(w, r)
		return
	}

	// 执行handler
	handler.ServeHTTP(w, r)
}

func NewServerMux() *Mux {
	return &Mux{
		Handles: map[string]http.Handler{},
		Filters: map[string]FilterFunc{},
	}
}

// DefaultHandle 定义默认的server mux 对象
var DefaultHandle = NewServerMux()

// HandleFunc 设置路由和对应的处理handler
func HandleFunc(pre string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultHandle.Handles[pre] = http.HandlerFunc(handler)
}

// HandleFunc 设置路由和对应的处理handler
func (mux *Mux) HandleFunc(pre string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Handles[pre] = http.HandlerFunc(handler)
}

func Handle(pre string, handler http.Handler) {
	DefaultHandle.Handles[pre] = handler
}

func (mux *Mux) Handle(pre string, handler http.Handler) {
	mux.Handles[pre] = handler
}

// FilterFunc 定义过滤器
// 参数一:过滤时所需要的参数
// 参数二:用于扩展所需的参数,或者或者返回值
type FilterFunc func(http.ResponseWriter, *http.Request) int

// Filter设置过滤器
func Filter(prefix string, filter FilterFunc) {
	DefaultHandle.Filters[prefix] = filter
}

// FilterFunc 设置过滤器
func (mux *Mux) FilterFunc(prefix string, filter FilterFunc) {
	mux.Filters[prefix] = filter
}

// deliverHandler 根据请求进来的路由找到对应的handler, 找不到则返回nil
func (mux *Mux) deliverHandler(path string) http.Handler {
	return mux.matchHandler(path)
}

// deliverFilter 根据请求进来的路由找到对应的过滤器,找不到则返回nil
func (mux *Mux) deliverFilter(path string) FilterFunc {
	for key, filter := range mux.Filters {
		if strings.HasPrefix(path, key) {
			return filter
		}
	}
	return nil
}

// Does path match pattern?
// copy from net/http/server.go
func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		// should not happen
		return false
	}
	n := len(pattern)
	if pattern[n-1] != '/' {
		return pattern == path
	}
	return len(path) >= n && path[0:n] == pattern
}

// Find a handler on a handler map given a path string
// Most-specific (longest) pattern wins
// copy from net/http/server.go and modify
func (mux *Mux) matchHandler(path string) (h http.Handler) {
	var n = 0
	for k, v := range mux.Handles {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v
		}
	}
	return
}

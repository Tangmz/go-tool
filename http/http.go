package http

import (
	"net/http"
	"strings"
)

const (
	// 过滤器错误码
	FILTER_ERROR = 1000
)

// http请求 继续往下执行标记码
const (
	REQUEST_CONTINUE int = 0 // 请求操作继续执行
	REQUEST_RETURN int = 1 // 请求操作停机执行
)

type Mux struct {
	Handles map[string]http.Handler // 请求路由
	Filter  map[string]FilterFunc   // 过滤器
}

// ServeHTTP 实现了http的ServeHTTP接口,以实现http的封装
func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// 根据请求的路由获取执行的handler
	handler := mux.deliverHandler(path)
	if handler == nil {
		http.NotFound(w, r)
		return
	}
	// 获取过滤器,如果有则执行
	filter := mux.deliverFilter(path)
	var filter_code int
	if filter != nil {
		filter_code = filter(w, r)
		if REQUEST_RETURN == filter_code {
			return
		}
	}
	// 执行handler
	handler.ServeHTTP(w, r)
}

func NewServerMux() *Mux {
	return &Mux{
		Handles:map[string]http.Handler{},
		Filter:map[string]FilterFunc{},
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
	mux.Handle(pre, handler)
}

// FilterFunc 定义过滤器
// 参数一:过滤时所需要的参数
// 参数二:用于扩展所需的参数,或者或者返回值
type FilterFunc func(http.ResponseWriter,*http.Request) int

// Filter设置过滤器
func Filter(prefix string, filter FilterFunc) {
	DefaultHandle.Filter[prefix] = filter
}

// Filter设置过滤器
func (mux *Mux) FilterFunc(prefix string, filter FilterFunc) {
	mux.Filter[prefix] = filter
}

// deliverHandler 根据请求进来的路由找到对应的handler, 找不到则返回nil
func (mux *Mux) deliverHandler(path string) http.Handler {
	for key, handler := range mux.Handles {
		if strings.HasPrefix(path, key) {
			return handler
		}
	}
	return nil
}

// deliverFilter 根据请求进来的路由找到对应的过滤器,找不到则返回nil
func (mux *Mux) deliverFilter(path string) FilterFunc {
	for key, filter := range mux.Filter {
		if strings.HasPrefix(path, key) {
			return filter
		}
	}
	return nil
}

func ListenAndServe(addr string, handler http.Handler) error {
	if handler == nil {
		handler = DefaultHandle
	}
	return http.ListenAndServe(addr, handler)
}

func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	if handler == nil {
		handler = DefaultHandle
	}
	return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
}
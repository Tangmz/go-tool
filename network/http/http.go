package http

import (
	"log"
	"net/http"
)

type ServeMux struct {
	m    map[string]muxEntry
	mAry []muxEntry
}

type muxEntry struct {
	h        http.Handler
	pattern  string
	explicit bool
}

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

func Handle(pattern string, handler http.Handler) {
	DefaultServeMux.Handle(pattern, handler)
}

func (mux *ServeMux) Handle(pattern string, handler http.Handler) {
	if pattern == "" {
		panic("http: invalid pattern " + pattern)
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if mux.m[pattern].explicit {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	me := muxEntry{h: handler, pattern: pattern}
	if _, ok := mux.m[pattern]; ok {
		for index, p := range mux.mAry {
			if pattern == p.pattern {
				mux.mAry[index] = me
			}
		}
	}
	mux.m[pattern] = me
}

func HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	mux.Handle(pattern, http.HandlerFunc(handler))
}

func (mux ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("http: ServeHTTP panic: %v\r\n", err)
		}
	}()

	//var path = r.URL.Path
}

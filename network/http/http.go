package http

import (
	"log"
	"net/http"
	"regexp"
)

type ServeMux struct {
	m        map[string]muxEntry
	mAry     []muxEntry
	filter   []muxEntry
	monitors []Monitor
}

type muxEntry struct {
	h        http.Handler
	pattern  string
	explicit bool
}

type Monitor interface {
	Start()
	Done()
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

func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("http: ServeHTTP panic: %v\r\n", err)
		}
	}()

	// start exec monitors
	var monitors = mux.monitors
	for _, monitor := range monitors {
		monitor.Start()
	}
	defer func() {
		for _, monitor := range monitors {
			monitor.Done()
		}
	}()

	// start exec filter

	// start exec handler
}

func (mux *ServeMux) Filter(w http.ResponseWriter, r *http.Request) {
	mux.match(w, r, "Filter")
}

func (mux *ServeMux) Handler(w http.ResponseWriter, r *http.Request) http.Handler {
	mux.match(w, r, "Mux")
}

func (mux *ServeMux) match(w http.ResponseWriter, r *http.Request, Type string) http.Handler {
	var handlers = []muxEntry{}
	if Type == "Filter" {
		handlers = mux.filter
	} else if Type == "Mux" {
		handlers = mux.mAry
	}
	for _, h := range handlers {
		if matched, _ := regexp.MatchString(r.URL.Path, h.pattern); matched {
			h.h.ServeHTTP(w, r)
		}
	}
	return nil
}

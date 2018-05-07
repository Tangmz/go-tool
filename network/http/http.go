package http

import (
	"net/http"
	"sync"
)

type ServeMux struct {
	middleWare  map[string]entity
	middleWares []entity
	handler     map[string]entity
	handlers    []entity

	lock *sync.Mutex
}

type entity struct {
	handler  http.Handler
	explicit bool
	pattern  string
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux { return new(ServeMux) }

var DefaultServeMux = NewServeMux()

func HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	mux.Handle(pattern, http.Handler(handler))
}

func Handle(pattern string, handler http.Handler) {
	DefaultServeMux.Handle(pattern, handler)
}

func (mux *ServeMux) Handle(pattern string, handler http.Handler) {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	if pattern == "" {
		panic("http: invalid pattern " + pattern)
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if mux.handler[pattern].explicit {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.handler == nil {
		mux.handler = make(map[string]entity)
	}

	mux.addHandler(pattern, handler)
}

func (mux *ServeMux) addHandler(pattern string, handler http.Handler) {
	var ent = entity{explicit: true, pattern: pattern, handler: handler}
	mux.handler[pattern] = ent
	for _, entry := range mux.handlers {
		if pattern == entry.pattern {
		}
	}
	mux.handlers = append(mux.handlers, ent)
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (s *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

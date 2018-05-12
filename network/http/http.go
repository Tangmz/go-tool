package http

import (
	"net/http"
	"path"
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
	mux.handlers = append(mux.handlers, ent)
}

func MiddleWareFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	DefaultServeMux.MiddleWare(pattern, http.Handler(handler))
}

func MiddleWare(pattern string, handler http.Handler) {
	DefaultServeMux.MiddleWare(pattern, handler)
}

func (mux *ServeMux) MiddleWare(pattern string, handler http.Handler) {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	if pattern == "" {
		panic("middleware: invalid pattern " + pattern)
	}
	if handler == nil {
		panic("middleware: nil handler")
	}
	if mux.handler[pattern].explicit {
		panic("middleware: multiple registrations for " + pattern)
	}

	if mux.handler == nil {
		mux.handler = make(map[string]entity)
	}

	mux.addMiddleWare(pattern, handler)
}

func (mux *ServeMux) addMiddleWare(pattern string, handler http.Handler) {
	var ent = entity{explicit: true, pattern: pattern, handler: handler}
	mux.middleWare[pattern] = ent
	mux.middleWares = append(mux.middleWares, ent)
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// find middleWare and exec it.
	for _, m := range mux.middleWares {
		path := cleanPath(r.URL.Path)
		if m.pattern == path {
			m.handler.ServeHTTP(w, r)
		}
	}

	// find http handler and exec it.
	for _, m := range mux.handlers {
		path := cleanPath(r.URL.Path)
		if m.pattern == path {
			m.handler.ServeHTTP(w, r)
		}
	}
}

// Return the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

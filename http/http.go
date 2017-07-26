package http

import "net/http"

type Mux struct {

}

func NewServerMux() *Mux {
	return &Mux{}
}

func HandleFunc(pre string, handler func(http.ResponseWriter, *http.Request)) {

}
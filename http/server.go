package http

import (
	"github.com/tangs-drm/go-tool/util"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
}

// WriteString write a string to the connection
func (w ResponseWriter) WriteString(data string) (int, error) {
	return w.Write([]byte(data))
}

// WriteValue write a value which will be trans to string to the connection
func (w ResponseWriter) WriteValue(data interface{}) (int, error) {
	var dataString = util.S2Json(data)
	return w.WriteString(dataString)
}

package http

import (
	"github.com/go-tool/util"
	"net/http"
)

type Response struct {
	Code    int      `json:"code"`
	Message string   `json:"message,omitempty"`
	Error   string   `json:"error,omitempty"`
	Data    util.Map `json:"data,omitempty"`
}

// WriteResponse write data to http.ResponseWriter
func WriteResponse(w http.ResponseWriter, code int, msg string, err error, data util.Map) (int, error) {
	var resp = Response{
		Code: code,
	}
	if len(msg) > 0 {
		resp.Message = msg
	}
	if data != nil {
		resp.Data = data
	}
	if err != nil {
		resp.Error = err.Error()
	}

	var result = util.S2Json(resp)
	return w.Write([]byte(result))
}

// WriteResponse write return valut to http.ResponseWriter if some error happened
func WriteResponseError(w http.ResponseWriter, code int, msg string, err error) (int, error) {
	return WriteResponse(w, code, msg, err, nil)
}

// WriteResponseSuccess write return value to http.ResponseWriter if http request succeed
func WriteResponseSuccess(w http.ResponseWriter, data util.Map) (int, error) {
	return WriteResponse(w, 0, "", nil, data)
}

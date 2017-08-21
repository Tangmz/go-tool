/*
	goalng http request tool.
	it contain func to request with method GET and POST.
*/
package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// httpRequest 发起一个http请求
// method 请求的方法, body 请求体，headers 设置的请求头， format， args两个合在一起拼接成一个url
func httpRequest(method string, body io.Reader, headers map[string]string, format string, args ...interface{}) ([]byte, error) {
	var client = http.Client{}
	var url = fmt.Sprintf(format, args...)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// HTTPGetString 开启了一个GET请求,获得字符串的返回值
func HTTPGetString(format string, args ...interface{}) (string, error) {
	bys, err := httpRequest("GET", nil, nil, format, args...)
	return string(bys), err
}

// HTTPGetMap 开启了一个GET请求,获得Map对象的返回值
func HTTPGetMap(format string, args ...interface{}) (Map, error) {
	bys, err := httpRequest("GET", nil, nil, format, args...)
	if err != nil {
		return nil, err
	}
	var result Map
	err = json.Unmarshal(bys, &result)
	return result, err
}

// HTTPPostString 开启了一个Post请求,获得字符串的返回值
func HTTPPostString(url string, headers map[string]string, body io.Reader) (string, error) {
	bys, err := httpRequest("POST", body, headers, "%v", url)
	return string(bys), err
}

// HTTPPostMap 开启了一个Post请求,获得Map对象的返回值
func HTTPPostMap(url string, headers map[string]string, body io.Reader) (Map, error) {
	bys, err := httpRequest("POST", body, headers, "%v", url)
	var result Map
	err = json.Unmarshal(bys, &result)
	return result, err
}

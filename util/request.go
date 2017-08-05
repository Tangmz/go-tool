/*
	goalng http request tool.
	it contain func to request with method GET and POST.
*/
package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
)

// HTTPGetString 开启了一个GET请求,获得字符串的返回值
func HTTPGetString(format string, args... interface{}) (string, error) {
	url := fmt.Sprintf(format, args...)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bys), nil
}

// HTTPGetMap 开启了一个GET请求,获得Map对象的返回值
func HTTPGetMap(format string, args... interface{}) (Map, error) {
	url := fmt.Sprintf(format, args...)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Map
	err = json.Unmarshal(bys, &result)
	return result, err
}

// HTTPPostString 开启了一个Post请求,获得字符串的返回值
func HTTPPostString(url string, contentType string, body io.Reader) (string, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bys), nil
}

// HTTPPostMap 开启了一个Post请求,获得Map对象的返回值
func HTTPPostMap(url string, contentType string, body io.Reader) (Map, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Map
	err = json.Unmarshal(bys, &result)
	return result, err
}
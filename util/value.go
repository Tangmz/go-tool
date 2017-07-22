package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
	"encoding/base64"
	"fmt"
	"crypto/rand"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UUIDString() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(GetMd5String(base64.URLEncoding.EncodeToString(b)))
}

func UUID() string {
	uuid := UUIDString()
	if "" == uuid {
		return ""
	}
	return fmt.Sprintf("%v-%v-%v-%v-%v", uuid[0:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
}
// TrimAryStringRepeat trim the s in every string in strs,
// repeat is decide result string array whether allowed repeat,
// false is don't allowed
func TrimAryStringRepeat(strs []string, s string, repeat bool) []string {
	var strMap = map[string]bool{}
	var results = []string{}
	for _, str := range strs {
		val := strings.Trim(str, s)
		if len(val) < 1 {
			continue
		}
		if strMap[val] && !repeat {
			continue
		}
		strMap[val] = true
		results = append(results, val)
	}
	return results
}

// TrimAryInt trim the same int value in vals or which value equal to arg
func TrimAryInt(vals []int, arg int) []int {
	var intMap = map[int]bool{}
	var result = []int{}
	for _, val := range vals {
		if intMap[val] {
			continue
		}
		intMap[val] = true
		result = append(result, val)
	}
	return result
}

// TrimAryInt trim the same int value int vals
func TrimAryIntRepeat(vals []int) []int {
	var intMap = map[int]bool{}
	var result = []int{}
	for _, val := range vals {
		if intMap[val] {
			continue
		}
		intMap[val] = true
		result = append(result, val)
	}
	return result
}
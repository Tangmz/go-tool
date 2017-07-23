package util

import (
	"encoding/json"
	"reflect"
)

type Map map[string]interface{}

// Exist judge if the key exist in map
func (m Map) Exist(key string) bool {
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

// String return the value type string by key
func (m Map) String(key string) string {
	var val string
	var ok bool
	if val, ok = m[key].(string); !ok {
		return ""
	}
	return val
}

// Int return the value type int by key
func (m Map) Int(key string) int {
	var val int
	var ok bool
	if val, ok = m[key].(int); !ok {
		return val
	}
	return val
}

// Int return the value type int32 by key
func (m Map) Int32(key string) int32 {
	var val int32
	var ok bool
	if val, ok = m[key].(int32); !ok {
		return val
	}
	return val
}

// Int return the value type int64 by key
func (m Map) Int64(key string) int64 {
	var val int64
	var ok bool
	if val, ok = m[key].(int64); !ok {
		return val
	}
	return val
}

// Int return the value type float32 by key
func (m Map) Float32(key string) float32 {
	var val float32
	var ok bool
	if val, ok = m[key].(float32); !ok {
		return val
	}
	return val
}

// Int return the value type float32 by key
func (m Map) Float64(key string) float64 {
	var val float64
	var ok bool
	if val, ok = m[key].(float64); !ok {
		return val
	}
	return val
}

// Int return the value type Map by key
func (m Map) Map(key string) Map {
	var val Map
	var ok bool
	if val, ok = m[key].(Map); !ok {
		return nil
	}
	return val
}

// S2Json trans data to json, e.g struct, map and so on.
func S2Json(data interface{}) string {
	bys, _ := json.Marshal(data)
	return string(bys)
}

// Json2S trans json to object
func Json2S(src string, dest interface{}) error {
	return json.Unmarshal([]byte(src), dest)
}

func TransType(val interface{}) interface{} {
	typ := reflect.TypeOf(val).Kind()
	switch typ {
	case reflect.Int:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
	case reflect.Array:
	}
	return typ
}
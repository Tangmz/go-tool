package util

import "testing"

func TestMap(t *testing.T) {
	m := Map{}

	// 检查语法错误
	m.String("halo")
	m.Int("halo")
	m.Int32("halo")
	m.Int64("halo")
	m.Float32("halo")
	m.Float64("halo")
	m.Map("halo")

	// 检查值
	m = Map{
		"string": "string",
		"int": 1,
		"int32": int32(2),
		"int64": int64(3),
		"float32": float32(1.1),
		"float64": float64(2.2),
		"Map": Map{"key": "val"},
	}
	if "string" != m.String("string") {
		t.Error(m.String("string"))
		return
	}
	if 1 != m.Int("int") {
		t.Error(m.Int("int"))
		return
	}
	if 2 != m.Int32("int32") {
		t.Error(m.Int32("int32"))
		return
	}
	if 3 != m.Int64("int64") {
		t.Error(m.Int64("int64"))
		return
	}
	if 1.1 != m.Float32("float32") {
		t.Error(m.Float32("float32"))
		return
	}
	if 2.2 != m.Float64("float64") {
		t.Error("float64")
		return
	}
	if nil == m.Map("Map") {
		t.Error(m.Map("Map"))
		return
	}
	mm := m.Map("Map")
	if "val" != mm.String("key") {
		t.Error(mm.String("key"))
		return
	}
	if true != m.Exist("string") {
		t.Error(m.Exist("string"))
	}
	if false != m.Exist("halo") {
		t.Error(m.Exist("halo"))
	}
}
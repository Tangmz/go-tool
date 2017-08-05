### golang-tool

该包包含了golang使用使用过程中的小工具函数


#### map

定义:

```golang
type Map map[string]interface{}
```

可以通过该对象转换value,见如下:
```
    m = Map{
        "string": "string",
        "int": "1",
        "int32": "2",
        "int64": "3",
        "float32": "1.1",
        "float64": 2.2,
        "Map": Map{"key": "val"},
    }

    m.Int("int) // 1 (type int)
    m.Int32("int64") // 3 (type int32)
    m.String("float64") // 2.2 (type string)
```

具体详情可见测试map_test.go
# log

创建一个简单的日志对象

用法如下:

```golang
package main

import (
	"github.com/tangs/log"
	"os"
	"fmt"
)

func main() {
	// 直接使用对象打印日志
	log.Debug("halo")

	// 手动设置打印日志到文件
	f, err := os.Create("log.log")
	if err != nil {
		log.Error("%v", err)
		return
	}
	log.Redirect(f)
	log.Debug("%v", "halo, log!")

	// 重定向日志到文件
	err = log.RedirectFile("log/log.log")
	if err != nil {
		fmt.Println("log ", err.Error())
		return
	}
	log.Debug("%v", "halo, RedirectFile")
}
```
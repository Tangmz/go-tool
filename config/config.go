package config

import (
	"github.com/tangs-drm/go-tool/util"
	"strings"
	"os"
	"bufio"
	"io"
	"fmt"
)

type Config struct {
	util.Map
}


// 配置文件的注释字符.如果遇到#,则忽略此行
const CONFIG_COMMENT = "#"
// 配置文件的配置项和值的连接符,如url=127.0.0.1
const CONFIG_EQUAL = "="
// 配置文件支持的文件后缀
const (
	CONFIG_CONF = ".conf"
)

func NewConfig() *Config {
	return &Config{
		Map: util.Map{},
	}
}

func checkExt(filename string) bool {
	return strings.HasSuffix(filename, CONFIG_CONF)
}

func (cf *Config) Config(filename string) error {
	if !checkExt(filename) {
		return util.Error("file type must be %v, invalid file name(%v)", CONFIG_CONF, filename)
	}
	return cf.ParseFile(filename)
}

func (cf *Config) ParseFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	var eof bool
	for {
		if eof {
			break
		}
		line, err := rd.ReadString('\n')
		if err != nil {
			if err.Error() != io.EOF.Error() {
				return err
			}
			eof = true
		}
		line = strings.TrimSpace(line)

		// 如果是空行,跳过
		if len(line) < 1 {
			continue
		}

		// 判断是否开头有注释符号
		if strings.HasPrefix(line, CONFIG_COMMENT) {
			continue
		}

		// 处理掉一行中的注释部分
		index := strings.Index(line, CONFIG_COMMENT)
		if index > -1 {
			line = line[0:index]
		}
		line = strings.TrimSpace(line)

		vals := strings.SplitN(line, CONFIG_EQUAL, 2)
		if len(vals) < 1 {
			continue
		}

		config_key := strings.TrimSpace(vals[0])
		config_val := ""
		if len(vals) == 2 {
			config_val = strings.TrimSpace(vals[1])
		}

		cf.Map[config_key] = config_val
	}
	return nil
}

func (cf *Config) Print() {
	for k, v := range cf.Map {
		fmt.Println(k, "---->", v)
	}
}
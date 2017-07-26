/*
	log 包,日志等级分为四个等级,DEBUG, INFO, WARNING, ERROR, 和通常的日志包套路一样,日志打印可以设置等级
	低于设置等级的日志会被过滤.
*/
package log

import (
	"log"
	"os"
	"fmt"
	"io"
	"path/filepath"
)

type logLevel int

type Logger struct {
	Level logLevel
	Log *log.Logger
}

const (
	DEBUG logLevel = iota
	INFO             = 2
	WARNING          = 3
	ERROR            = 4
)

// 包内私有对象
var logger = NewLogger()

// NewLogger返回一个新log对象
func NewLogger() *Logger {
	return &Logger{
		Log: log.New(os.Stdout, "",  log.LstdFlags|log.Lshortfile),
	}
}

// SetLevel设置日子的打印等级
func (l *Logger) SetLevel(level int) {
	l.Level = logLevel(level)
}

// Redirect日志重定向,可以指定到控制台,文件等
func Redirect(writer io.Writer) {
	logger.Log.SetOutput(writer)
}

// RedirectFile日志重定向到文件,可以创建多层文件路径,如 /var/log/server/log/log.log
func RedirectFile(file string) error {
	var err error
	fp := filepath.Dir(file)
	err = os.MkdirAll(fp, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0755)
	if err != nil {
		return err
	}
	Redirect(f)
	return nil
}

// Debug 信息
func Debug(format string, args... interface{}) {
	if DEBUG < logger.Level {
		return
	}
	logger.Log.Output(2,  fmt.Sprintf("[D] " + format, args...))
}

// Info 信息
func Info(format string, args... interface{})  {
	if INFO < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[I] " + format, args...))
}

// Warning 信息
func Warning(format string, args... interface{})  {
	if WARNING < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[W] " + format, args...))
}

// Error 信息
func Error(format string, args... interface{})  {
	if ERROR < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[E] " + format, args...))
}

// Debug 信息
func (l *Logger) Debug(format string, args... interface{}) {
	if DEBUG < logger.Level {
		return
	}
	logger.Log.Output(2,  fmt.Sprintf("[D] " + format, args...))
}

// Info 信息
func (l *Logger) Info(format string, args... interface{})  {
	if INFO < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[I] " + format, args...))
}

// Warning 信息
func (l *Logger) Warning(format string, args... interface{})  {
	if WARNING < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[W] " + format, args...))
}

// Error 信息
func (l *Logger) Error(format string, args... interface{})  {
	if ERROR < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[E] " + format, args...))
}

// 自定义callpath的 Debug信息,常用于测试
func LogD_(callpath int, format string, args... interface{}) {
	logger.Log.Output(callpath, fmt.Sprintf("[D] " + format, args...))
}
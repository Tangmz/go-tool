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

var logger = NewLogger()

func NewLogger() *Logger {
	return &Logger{
		Log: log.New(os.Stdout, "",  log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) SetLevel(level int) {
	l.Level = logLevel(level)
}

func Redirect(writer io.Writer) {
	logger.Log.SetOutput(writer)
}

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

func Debug(format string, args... interface{}) {
	if DEBUG < logger.Level {
		return
	}
	logger.Log.Output(2,  fmt.Sprintf("[D] " + format, args...))
}

func Info(format string, args... interface{})  {
	if INFO < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[I] " + format, args...))
}

func Warning(format string, args... interface{})  {
	if WARNING < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[W] " + format, args...))
}

func Error(format string, args... interface{})  {
	if ERROR < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[E] " + format, args...))
}


func (l *Logger) Debug(format string, args... interface{}) {
	if DEBUG < logger.Level {
		return
	}
	logger.Log.Output(2,  fmt.Sprintf("[D] " + format, args...))
}

func (l *Logger) Info(format string, args... interface{})  {
	if INFO < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[I] " + format, args...))
}

func (l *Logger) Warning(format string, args... interface{})  {
	if WARNING < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[W] " + format, args...))
}

func (l *Logger) Error(format string, args... interface{})  {
	if ERROR < logger.Level {
		return
	}
	logger.Log.Output(2, fmt.Sprintf("[E] " + format, args...))
}
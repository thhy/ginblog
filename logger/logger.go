package logger

import (
	"fmt"
	"runtime"
)

const (
	DEBUG = iota
	INFO
	ERROR
	FATAL
)

var LOGGERLEVEL = DEBUG

var LEVELINFO = []string{
	"DEBUG",
	"INFO",
	"ERROR",
	"FATAL",
}

func Log(level int, str ...interface{}) {
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		if level >= LOGGERLEVEL {
			fmt.Printf("[%s: %d][func:%s][%s]%+v \n", file, line, runtime.FuncForPC(funcName).Name(), LEVELINFO[level], str)
		}
	}
}

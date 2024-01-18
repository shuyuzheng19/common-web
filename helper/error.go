package helper

import (
	"errors"
	"log"
)

// ErrorPanic 全局error处理
func ErrorPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

// ErrorPanicAndMessage 全局error处理
func ErrorPanicAndMessage(err error, message string) {
	if err != nil {
		log.Panicln(errors.New(message))
	}
}

// ErrorCommonF  全局error处理
func ErrorCommonF(f interface{}) {
	panic(f)
}

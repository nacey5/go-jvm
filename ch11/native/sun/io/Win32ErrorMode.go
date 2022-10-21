package io

import (
	"go-jvm/ch11/native"
	"go-jvm/ch11/runtime_data_area"
)

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *runtime_data_area.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}

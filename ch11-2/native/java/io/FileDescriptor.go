package io

import "go-jvm/ch11-2/native"
import "go-jvm/ch11-2/runtime_data_area"

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
}

// private static native long set(int d);
// (I)J
func set(frame *runtime_data_area.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}

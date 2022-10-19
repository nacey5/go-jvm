package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
)

func init() {
	native.Registry("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// public static Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *runtime_data_area.Frame) {
	//todo 在后面实现
}

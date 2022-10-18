package misc

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
	"go-jvm/ch09/runtime_data_area/heap"
)

func init() {
	native.Registry("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize()
func initialize(frame *runtime_data_area.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")
	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)

	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)
}

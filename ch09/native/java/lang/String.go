package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
	"go-jvm/ch09/runtime_data_area/heap"
)

func init() {
	native.Registry("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
func intern(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}

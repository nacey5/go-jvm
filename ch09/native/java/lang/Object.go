package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
)

func init() {
	native.Registry("java/lang/Object", "getClass", "()Ljava/lang/Class", getClass)
}

// public final native Class<?> getClass();
func getClass(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
	"unsafe"
)

func init() {
	native.Registry("java/lang/Object", "getClass", "()Ljava/lang/Class", getClass)
	native.Registry("java/lang/Object", "hashCode", "I()", hashCode)
	native.Registry("java/lang.Object", "clone", "()Ljava/lang/Object;", clone)
}

func clone(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}

func hashCode(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// public final native Class<?> getClass();
func getClass(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

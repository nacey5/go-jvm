package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
	"go-jvm/ch09/runtime_data_area/heap"
)

const jlClass = "java/lang/Class"

func init() {
	native.Registry(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Registry(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Registry(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

func desiredAssertionStatus0(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushBoolean(false)
}

// private native String getName0();
func getName0(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObject := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObject)
}

// static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *runtime_data_area.Frame) {
	nameObject := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObject)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}

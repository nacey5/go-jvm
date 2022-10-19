package lang

import "go-jvm/ch10/native"
import "go-jvm/ch10/runtime_data_area"
import "go-jvm/ch10/runtime_data_area/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}

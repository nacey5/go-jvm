package reserved

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/native"
	_ "go-jvm/ch09/native/java/lang"
	"go-jvm/ch09/runtime_data_area"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (this *INVOKE_NATIVE) Execute(frame *runtime_data_area.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnstatisfiedLinkError" + methodInfo)
	}
	nativeMethod(frame)
}

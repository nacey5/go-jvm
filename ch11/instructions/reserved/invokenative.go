package reserved

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"
import "go-jvm/ch11/native"
import _ "go-jvm/ch11/native/java/lang"
import _ "go-jvm/ch11/native/sun/misc"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (this *INVOKE_NATIVE) Execute(frame *runtime_data_area.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}

package reserved

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"
import "go-jvm/ch11-2/native"
import _ "go-jvm/ch11-2/native/java/io"
import _ "go-jvm/ch11-2/native/java/lang"
import _ "go-jvm/ch11-2/native/java/security"
import _ "go-jvm/ch11-2/native/java/util/concurrent/atomic"
import _ "go-jvm/ch11-2/native/sun/io"
import _ "go-jvm/ch11-2/native/sun/misc"
import _ "go-jvm/ch11-2/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *runtime_data_area.Frame) {
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

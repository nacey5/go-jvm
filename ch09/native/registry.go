package native

import "go-jvm/ch09/runtime_data_area"

// 定义一个本地方法接口
type NativeMethod func(frame *runtime_data_area.Frame)

func emptyNativeMethod(frame *runtime_data_area.Frame) {
	//do nothing
}

// 本地方法注册表
var registry = map[string]NativeMethod{}

func Registry(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

package references

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
	"go-jvm/ch08/runtime-data-area/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

// 解析静态方法
func (this *INVOKE_STATIC) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(this.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPc()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, resolvedMethod)
}

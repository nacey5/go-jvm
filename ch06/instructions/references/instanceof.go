package references

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
	"go-jvm/ch06/runtime-data-area/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

// 判断引用类型，如果obj==null的时候，if全为false
func (this *INSTANCE_OF) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

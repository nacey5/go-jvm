package references

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"
import "go-jvm/ch09-2/rtda/heap"

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (this *INSTANCE_OF) Execute(frame *rtda.Frame) {
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

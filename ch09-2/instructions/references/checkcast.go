package references

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"
import "go-jvm/ch09-2/rtda/heap"

// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }

func (this *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}

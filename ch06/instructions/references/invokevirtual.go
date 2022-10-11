package references

import (
	"fmt"
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
	"go-jvm/ch06/runtime-data-area/heap"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_VIRTUAL) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V":
			fmt.Printf("%c\n", stack.PopInt())
		case "(I)V", "(B)V", "(S)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(F)V":
			fmt.Printf("%v\n", stack.PopFloat())
		case "(J)V":
			fmt.Printf("%v\n", stack.PopLong())
		case "(D)V":
			fmt.Printf("%v\n", stack.PopDouble())
		default:
			panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}

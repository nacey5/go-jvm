package math

import "go-jvm/ch10/instructions/base"
import "go-jvm/ch10/runtime_data_area"

// Boolean XOR int
type IXOR struct{ base.NoOperandsInstruction }

func (this *IXOR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct{ base.NoOperandsInstruction }

func (this *LXOR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}

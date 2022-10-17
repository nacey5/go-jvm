package math

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

// Boolean OR int
type IOR struct{ base.NoOperandsInstruction }

func (this *IOR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperandsInstruction }

func (this *LOR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}

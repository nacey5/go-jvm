package stack

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

func (this *POP) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

func (this *POP2) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

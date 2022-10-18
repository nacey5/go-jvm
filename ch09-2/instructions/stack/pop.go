package stack

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

func (this *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

func (this *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

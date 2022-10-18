package stack

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (this *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

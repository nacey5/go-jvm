package stack

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (this *SWAP) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

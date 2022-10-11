package stack

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

type SWAP struct {
	base.NoOperandsInstruction
}

// swap指令交换栈顶的两个变量
func (this *SWAP) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

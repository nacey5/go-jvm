package stack

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}

// pop指令把栈顶变量弹出
func (this *POP) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// double和long类型占据两个位置
func (this *POP2) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

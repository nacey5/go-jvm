package math

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

// 布尔运算指令
type IAND struct {
	base.NoOperandsInstruction
}

type LAND struct {
	base.NoOperandsInstruction
}

// 与运算
func (this *IAND) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (this *LAND) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}

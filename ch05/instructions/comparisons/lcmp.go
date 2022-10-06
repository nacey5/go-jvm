package comparisons

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

// LCMP 比较long类型
type LCMP struct {
	base.NoOperandsInstruction
}

func (this *LCMP) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

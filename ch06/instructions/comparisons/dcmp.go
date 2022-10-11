package comparisons

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

// Compare double
type DCMPG struct{ base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *runtime_data_area.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (self *DCMPL) Execute(frame *runtime_data_area.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *runtime_data_area.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

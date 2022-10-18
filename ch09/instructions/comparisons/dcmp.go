package comparisons

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Compare double
type DCMPG struct{ base.NoOperandsInstruction }

func (this *DCMPG) Execute(frame *runtime_data_area.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (this *DCMPL) Execute(frame *runtime_data_area.Frame) {
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

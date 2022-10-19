package comparisons

import "go-jvm/ch10/instructions/base"
import "go-jvm/ch10/runtime_data_area"

// Compare float
type FCMPG struct{ base.NoOperandsInstruction }

func (this *FCMPG) Execute(frame *runtime_data_area.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (this *FCMPL) Execute(frame *runtime_data_area.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *runtime_data_area.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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

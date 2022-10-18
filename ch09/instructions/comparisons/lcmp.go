package comparisons

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Compare long
type LCMP struct{ base.NoOperandsInstruction }

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

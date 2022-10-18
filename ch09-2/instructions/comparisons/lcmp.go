package comparisons

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Compare long
type LCMP struct{ base.NoOperandsInstruction }

func (this *LCMP) Execute(frame *rtda.Frame) {
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

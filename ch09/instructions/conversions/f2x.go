package conversions

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Convert float to double
type F2D struct{ base.NoOperandsInstruction }

func (this *F2D) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct{ base.NoOperandsInstruction }

func (this *F2I) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct{ base.NoOperandsInstruction }

func (this *F2L) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}

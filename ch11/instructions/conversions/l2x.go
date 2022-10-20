package conversions

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Convert long to double
type L2D struct{ base.NoOperandsInstruction }

func (this *L2D) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct{ base.NoOperandsInstruction }

func (this *L2F) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct{ base.NoOperandsInstruction }

func (this *L2I) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

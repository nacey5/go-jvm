package conversions

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (this *D2F) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (this *D2I) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (this *D2L) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}

package conversions

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

type F2D struct{ base.NoOperandsInstruction }

func (self *F2D) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

type F2I struct{ base.NoOperandsInstruction }

func (self *F2I) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

type F2L struct{ base.NoOperandsInstruction }

func (self *F2L) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}

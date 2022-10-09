package conversions

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

// 类型转换指令
type D2F struct {
	base.NoOperandsInstruction
}
type D2I struct {
	base.NoOperandsInstruction
}
type D2L struct {
	base.NoOperandsInstruction
}

func (this *D2I) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
func (this *D2F) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}
func (this *D2L) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}

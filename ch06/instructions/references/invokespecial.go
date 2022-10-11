package references

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

// hack!
func (this *INVOKE_SPECIAL) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PopRef()
}

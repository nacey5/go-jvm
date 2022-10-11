package constants

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

// nop指令
type NOP struct {
	base.NoOperandsInstruction
}

func (this *NOP) Execute(frame *runtime_data_area.Frame) {
	//nothing to do
}

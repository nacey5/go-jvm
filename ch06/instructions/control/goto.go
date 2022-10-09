package control

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

type GOTO struct {
	base.BranchInstruction
}

func (this *GOTO) Execute(frame *runtime_data_area.Frame) {
	base.Branch(frame, this.Offset)
}

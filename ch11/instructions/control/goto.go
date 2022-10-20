package control

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (this *GOTO) Execute(frame *runtime_data_area.Frame) {
	base.Branch(frame, this.Offset)
}

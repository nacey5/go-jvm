package control

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *runtime_data_area.Frame) {
	base.Branch(frame, self.Offset)
}

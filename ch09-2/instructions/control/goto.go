package control

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (this *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.Offset)
}

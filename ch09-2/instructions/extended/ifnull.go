package extended

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (this *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, this.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (this *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, this.Offset)
	}
}

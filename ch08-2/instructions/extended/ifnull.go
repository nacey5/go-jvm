package extended

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (this *IFNULL) Execute(frame *runtime_data_area.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, this.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (this *IFNONNULL) Execute(frame *runtime_data_area.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, this.Offset)
	}
}

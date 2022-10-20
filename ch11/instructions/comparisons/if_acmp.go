package comparisons

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (this *IF_ACMPEQ) Execute(frame *runtime_data_area.Frame) {
	if _acmp(frame) {
		base.Branch(frame, this.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (this *IF_ACMPNE) Execute(frame *runtime_data_area.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, this.Offset)
	}
}

func _acmp(frame *runtime_data_area.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}

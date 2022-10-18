package comparisons

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (this *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, this.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (this *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, this.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}

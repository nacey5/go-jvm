package comparisons

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (this *IF_ACMPEQ) Execute(frame *runtime_data_area.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, this.Offset)
	}
}

func (this *IF_ACMPNE) Execute(frame *runtime_data_area.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, this.Offset)
	}
}

func _acmp(frame *runtime_data_area.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}

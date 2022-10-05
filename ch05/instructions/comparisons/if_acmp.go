package comparisons

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

type IF_ACMPEQ struct {
	base.NoOperandsInstruction
}

type IF_ACMPNE struct {
	base.NoOperandsInstruction
}

func (this *IF_ACMPEQ) Execute(frame runtime_data_area.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *runtime_data_area.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *runtime_data_area.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}

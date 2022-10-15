package comparisons

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

type IFEQ struct {
	base.BranchInstruction
}
type IFNE struct {
	base.BranchInstruction
}

type IFLT struct {
	base.BranchInstruction
}

type IFLE struct {
	base.BranchInstruction
}

type IFGT struct {
	base.BranchInstruction
}

type IFGE struct {
	base.BranchInstruction
}

func (this *IFEQ) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, this.Offset)
	}
}
func (this *IFNE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, this.Offset)
	}
}
func (this *IFLT) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, this.Offset)
	}
}
func (this *IFLE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, this.Offset)
	}
}
func (this *IFGT) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, this.Offset)
	}
}
func (this *IFGE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, this.Offset)
	}
}

package comparisons

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (this *IFEQ) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (this *IFNE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (this *IFLT) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (this *IFLE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (this *IFGT) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (this *IFGE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, this.Offset)
	}
}

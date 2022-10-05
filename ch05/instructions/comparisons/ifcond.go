package comparisons

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

type IFEQ struct {
	base.NoOperandsInstruction
}
type IFNE struct {
	base.NoOperandsInstruction
}

type IFLT struct {
	base.NoOperandsInstruction
}

type IFLE struct {
	base.NoOperandsInstruction
}

type IFGT struct {
	base.NoOperandsInstruction
}

type IFGE struct {
	base.NoOperandsInstruction
}

func (this *IFEQ) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, this.Offset)
	}
}
func (self *IFNE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFLT) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFLE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFGT) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFGE) Execute(frame *runtime_data_area.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}

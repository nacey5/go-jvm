package comparisons

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

type IF_ICMPEQ struct {
	base.NoOperandsInstruction
}
type IF_ICMPNE struct {
	base.NoOperandsInstruction
}
type IF_ICMPLT struct {
	base.NoOperandsInstruction
}
type IF_ICMPLE struct {
	base.NoOperandsInstruction
}
type IF_ICMPGT struct {
	base.NoOperandsInstruction
}
type IF_ICMPGE struct {
	base.NoOperandsInstruction
}

func (self *IF_ICMPEQ) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

func (this *IF_ICMPNE) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, this.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}

func _icmpPop(frame *runtime_data_area.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}

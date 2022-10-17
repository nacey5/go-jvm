package comparisons

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

func (this *IF_ICMPEQ) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (this *IF_ICMPNE) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (this *IF_ICMPLT) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (this *IF_ICMPLE) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (this *IF_ICMPGT) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (this *IF_ICMPGE) Execute(frame *runtime_data_area.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, this.Offset)
	}
}

func _icmpPop(frame *runtime_data_area.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}

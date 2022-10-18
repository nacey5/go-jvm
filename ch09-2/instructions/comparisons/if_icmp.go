package comparisons

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

func (this *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (this *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (this *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (this *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (this *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (this *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, this.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}

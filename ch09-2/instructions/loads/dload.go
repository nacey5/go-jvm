package loads

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Load double from local variable
type DLOAD struct{ base.Index8Instruction }

func (this *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, this.Index)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (this *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (this *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (this *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (this *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

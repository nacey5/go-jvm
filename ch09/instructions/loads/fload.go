package loads

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Load float from local variable
type FLOAD struct{ base.Index8Instruction }

func (this *FLOAD) Execute(frame *runtime_data_area.Frame) {
	_fload(frame, this.Index)
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

func (this *FLOAD_0) Execute(frame *runtime_data_area.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

func (this *FLOAD_1) Execute(frame *runtime_data_area.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

func (this *FLOAD_2) Execute(frame *runtime_data_area.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

func (this *FLOAD_3) Execute(frame *runtime_data_area.Frame) {
	_fload(frame, 3)
}

func _fload(frame *runtime_data_area.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

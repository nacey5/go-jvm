package loads

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

// Load long from local variable
type LLOAD struct{ base.Index8Instruction }

func (this *LLOAD) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, this.Index)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (this *LLOAD_0) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (this *LLOAD_1) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (this *LLOAD_2) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (this *LLOAD_3) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 3)
}

func _lload(frame *runtime_data_area.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

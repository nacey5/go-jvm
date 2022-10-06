package loads

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

// Load reference from local variable
type ALOAD struct{ base.Index8Instruction }

func (this *ALOAD) Execute(frame *runtime_data_area.Frame) {
	_aload(frame, this.Index)
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (this *ALOAD_0) Execute(frame *runtime_data_area.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (this *ALOAD_1) Execute(frame *runtime_data_area.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (this *ALOAD_2) Execute(frame *runtime_data_area.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (this *ALOAD_3) Execute(frame *runtime_data_area.Frame) {
	_aload(frame, 3)
}

func _aload(frame *runtime_data_area.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

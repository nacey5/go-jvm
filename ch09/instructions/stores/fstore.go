package stores

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Store float into local variable
type FSTORE struct{ base.Index8Instruction }

func (this *FSTORE) Execute(frame *runtime_data_area.Frame) {
	_fstore(frame, uint(this.Index))
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (this *FSTORE_0) Execute(frame *runtime_data_area.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (this *FSTORE_1) Execute(frame *runtime_data_area.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (this *FSTORE_2) Execute(frame *runtime_data_area.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (this *FSTORE_3) Execute(frame *runtime_data_area.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *runtime_data_area.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

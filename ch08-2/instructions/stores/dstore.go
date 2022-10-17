package stores

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

// Store double into local variable
type DSTORE struct{ base.Index8Instruction }

func (this *DSTORE) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, uint(this.Index))
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (this *DSTORE_0) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (this *DSTORE_1) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (this *DSTORE_2) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (this *DSTORE_3) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *runtime_data_area.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

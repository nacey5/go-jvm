package stores

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

// Store double into local variable
type DSTORE struct{ base.Index8Instruction }

func (self *DSTORE) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, uint(self.Index))
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (self *DSTORE_0) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (self *DSTORE_1) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (self *DSTORE_2) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (self *DSTORE_3) Execute(frame *runtime_data_area.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *runtime_data_area.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

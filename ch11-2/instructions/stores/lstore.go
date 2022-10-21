package stores

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"

// Store long into local variable
type LSTORE struct{ base.Index8Instruction }

func (self *LSTORE) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE_0 struct{ base.NoOperandsInstruction }

func (self *LSTORE_0) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }

func (self *LSTORE_1) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }

func (self *LSTORE_2) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }

func (self *LSTORE_3) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *runtime_data_area.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

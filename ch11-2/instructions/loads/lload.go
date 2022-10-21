package loads

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"

// Load long from local variable
type LLOAD struct{ base.Index8Instruction }

func (self *LLOAD) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, self.Index)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *runtime_data_area.Frame) {
	_lload(frame, 3)
}

func _lload(frame *runtime_data_area.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

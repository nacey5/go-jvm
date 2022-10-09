package stores

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

type LSTORE struct {
	base.Index8Instruction
}
type LSTORE_0 struct {
	base.NoOperandsInstruction
}
type LSTORE_1 struct {
	base.NoOperandsInstruction
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

// 公共方法
func _lstore(frame *runtime_data_area.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (this *LSTORE) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, uint(this.Index))
}
func (this *LSTORE_0) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 0)
}
func (this *LSTORE_1) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 1)
}
func (this *LSTORE_2) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 2)
}
func (this *LSTORE_3) Execute(frame *runtime_data_area.Frame) {
	_lstore(frame, 3)
}

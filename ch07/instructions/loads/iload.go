package loads

import (
	"go-jvm/ch07/instructions/base"
	runtime_data_area "go-jvm/ch07/runtime-data-area"
)

type ILOAD struct {
	base.Index8Instruction
}
type ILOAD_0 struct {
	base.NoOperandsInstruction
}
type ILOAD_1 struct {
	base.NoOperandsInstruction
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

// 避免重复代码，定义一个函数提供iload系列指令使用
func _iload(frame *runtime_data_area.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

// iload的指令索引来自操作数
func (this *ILOAD) Execute(frame *runtime_data_area.Frame) {
	_iload(frame, uint(this.Index))
}
func (this *ILOAD_0) Execute(frame *runtime_data_area.Frame) {
	_iload(frame, 0)
}
func (this *ILOAD_1) Execute(frame *runtime_data_area.Frame) {
	_iload(frame, 1)
}
func (this *ILOAD_2) Execute(frame *runtime_data_area.Frame) {
	_iload(frame, 2)
}
func (this *ILOAD_3) Execute(frame *runtime_data_area.Frame) {
	_iload(frame, 3)
}

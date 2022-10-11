package constants

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

// BIPUSH push byte
type BIPUSH struct {
	val int8
}

// SIPUSH push short
type SIPUSH struct {
	val int16
}

func (this *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}

func (this *BIPUSH) Execute(frame *runtime_data_area.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

func (this *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}

func (this *SIPUSH) Execute(frame *runtime_data_area.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

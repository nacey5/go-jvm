package constants

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Push byte
type BIPUSH struct {
	val int8
}

func (this *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}
func (this *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type SIPUSH struct {
	val int16
}

func (this *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}
func (this *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

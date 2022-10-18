package math

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (this *IINC) FetchOperands(reader *base.BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
	this.Const = int32(reader.ReadInt8())
}

func (this *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(this.Index)
	val += this.Const
	localVars.SetInt(this.Index, val)
}

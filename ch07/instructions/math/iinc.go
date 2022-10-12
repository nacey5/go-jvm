package math

import (
	"go-jvm/ch07/instructions/base"
	runtime_data_area "go-jvm/ch07/runtime-data-area"
)

//给局部变量表中的int类型增加常量值

type IINC struct {
	Index uint
	Const int32
}

// 从字节码中读取操作数
func (this *IINC) FetchOperands(reader *base.BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
	this.Const = int32(reader.ReadInt8())
}

// 从局部变量表中读取变量，给他加上常量值，再写回
func (this *IINC) Execute(frame *runtime_data_area.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(this.Index)
	val += this.Const
	localVars.SetInt(this.Index, val)
}

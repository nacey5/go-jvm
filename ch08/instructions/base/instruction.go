package base

import runtime_data_area "go-jvm/ch08/runtime-data-area"

// Instruction 指令接口
type Instruction interface {
	// FetchOperands 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// Execute 执行指令逻辑
	Execute(frame *runtime_data_area.Frame)
}

// 空指令
type NoOperandsInstruction struct {
	//nothing to do /nop
}

func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

// 表示跳转指令
type BranchInstruction struct {
	//跳转偏移量
	Offset int
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = int(reader.ReadInt16())
}

// Index8Instruction 存储和加载类指令需要根据索引存取局部变量表
type Index8Instruction struct {
	//局部变量表索引
	Index uint
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
}

// Index16Instruction 有些指令需要访问运行时常量池，常量池索引由两个字节给出
type Index16Instruction struct {
	//表示常量池索引
	Index uint
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadUint16())
}

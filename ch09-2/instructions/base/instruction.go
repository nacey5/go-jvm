package base

import "go-jvm/ch09-2/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
	// empty
}

func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadUint16())
}

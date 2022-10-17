package loads

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
	"go-jvm/ch08/runtime-data-area/heap"
)

type AALOAD struct {
	base.NoOperandsInstruction
}
type BALOAD struct {
	base.NoOperandsInstruction
}
type CALOAD struct {
	base.NoOperandsInstruction
}
type DALOAD struct {
	base.NoOperandsInstruction
}
type FALOAD struct {
	base.NoOperandsInstruction
}
type IALOAD struct {
	base.NoOperandsInstruction
}
type LALOAD struct {
	base.NoOperandsInstruction
}
type SALOAD struct {
	base.NoOperandsInstruction
}

func (this *AALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	CheckNotNil(arrRef)
	refs := arrRef.Refs()
	CheckIndex(len(refs), index)
	stack.PushRef(refs[index])
}

func CheckIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("java.lang.NullPointerException")
	}
}

// 检查
func CheckNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func (this *BALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

func (this *CALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

func (this *DALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

func (this *FALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

func (this *IALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

func (this *LALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

func (this *SALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

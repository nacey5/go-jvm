package loads

import "go-jvm/ch10/instructions/base"
import "go-jvm/ch10/runtime_data_area"
import "go-jvm/ch10/runtime_data_area/heap"

// Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

func (this *AALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }

func (this *BALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

// Load char from array
type CALOAD struct{ base.NoOperandsInstruction }

func (this *CALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// Load double from array
type DALOAD struct{ base.NoOperandsInstruction }

func (this *DALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

// Load float from array
type FALOAD struct{ base.NoOperandsInstruction }

func (this *FALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

// Load int from array
type IALOAD struct{ base.NoOperandsInstruction }

func (this *IALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

// Load long from array
type LALOAD struct{ base.NoOperandsInstruction }

func (this *LALOAD) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

// Load short from array
type SALOAD struct{ base.NoOperandsInstruction }

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

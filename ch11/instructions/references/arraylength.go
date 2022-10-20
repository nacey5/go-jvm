package references

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (this *ARRAY_LENGTH) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}

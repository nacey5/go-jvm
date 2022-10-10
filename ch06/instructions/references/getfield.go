package references

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
	"go-jvm/ch06/runtime-data-area/heap"
)

type GET_FIELD struct {
	base.Index16Instruction
}

func (this *GET_FIELD) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}

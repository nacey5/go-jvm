package constants

import (
	"go-jvm/ch06/instructions/base"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

// 加载指令
type LDC struct {
	base.Index8Instruction
}
type LDC_W struct {
	base.Index16Instruction
}
type LDC2_W struct {
	base.Index16Instruction
}

func (this *LDC) Execute(frame *runtime_data_area.Frame) {
	_ldc(frame, this.Index)
}

func (this *LDC_W) Execute(frame *runtime_data_area.Frame) {
	_ldc(frame, this.Index)
}

func _ldc(frame *runtime_data_area.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// todo 待完成的加载指令
	// case string:
	// case *heap.ClassRef
	default:
		panic("todo :ldc!")
	}
}

func (this *LDC2_W) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(this.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

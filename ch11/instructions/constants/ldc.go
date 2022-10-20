package constants

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"
import "go-jvm/ch11/runtime_data_area/heap"

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

func (this *LDC) Execute(frame *runtime_data_area.Frame) {
	_ldc(frame, this.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (this *LDC_W) Execute(frame *runtime_data_area.Frame) {
	_ldc(frame, this.Index)
}

func _ldc(frame *runtime_data_area.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

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

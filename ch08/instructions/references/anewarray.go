package references

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
	"go-jvm/ch08/runtime-data-area/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (this *ANEW_ARRAY) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

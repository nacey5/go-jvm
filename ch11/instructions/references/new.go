package references

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"
import "go-jvm/ch11/runtime_data_area/heap"

// Create new object
type NEW struct{ base.Index16Instruction }

func (this *NEW) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

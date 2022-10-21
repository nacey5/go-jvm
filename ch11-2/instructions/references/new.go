package references

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"
import "go-jvm/ch11-2/runtime_data_area/heap"

// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
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

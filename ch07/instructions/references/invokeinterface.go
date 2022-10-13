package references

import (
	"go-jvm/ch07/instructions/base"
	runtime_data_area "go-jvm/ch07/runtime-data-area"
	"go-jvm/ch07/runtime-data-area/heap"
)

// count 表示了slot数，zero进行补位以满足2^n
type INVOKE_INTERFACE struct {
	index uint
	//count uint8
	//zero uint8
}

func (this *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	this.index = uint(reader.ReadUint16())
	reader.ReadUint8() //count
	reader.ReadUint8() // must be zero
}

func (this *INVOKE_INTERFACE) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(this.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolveInterfaceMethod()
	//如果拿到的方法是私有或者静态的
	if resolvedMethod.IsStrict() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//从操作数栈中弹出引用，如果位nil或者并没有实现该方法所需要的接口
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//调用最终要查找的方法，如果不是public或者是抽象方法
	methodTobeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodTobeInvoked == nil || methodTobeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodTobeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	//一切正常
	base.InvokeMethod(frame, methodTobeInvoked)
}

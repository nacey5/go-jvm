package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
	"go-jvm/ch09/runtime_data_area/heap"
)

func init() {
	native.Registry("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arrayCopy)
}

// 数组复制
func arrayCopy(frame *runtime_data_area.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

// 检查数组是否匹配
func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}

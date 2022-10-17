package references

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
	"go-jvm/ch08/runtime-data-area/heap"
)

type MULTI_ANEW_ARRAY struct {
	//索引
	index uint16
	//索引所对应的引用
	dimensions uint8
}

func (this *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	this.index = reader.ReadUint16()
	this.dimensions = reader.ReadUint8()
}

func (this *MULTI_ANEW_ARRAY) Execute(frame *runtime_data_area.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(this.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	//必须保证多维数组弹出的n个int值确保他们都大于等于0
	counts := popAndCheckCounts(stack, int(this.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}
	return arr
}

func popAndCheckCounts(stack *runtime_data_area.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

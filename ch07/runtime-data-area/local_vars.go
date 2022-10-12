package runtime_data_area

import (
	"go-jvm/ch07/runtime-data-area/heap"
	"math"
)

type LocalVars []Slot

// 创建新的局部变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// SetInt 下列方法都没有对boolean,byte,short,char类型存取，所有都可以转成int进行处理
// 对Int类型的操作
func (this LocalVars) SetInt(index uint, val int32) {
	this[index].num = val
}

func (this LocalVars) GetInt(index uint) int32 {
	return this[index].num
}

// SetFloat 对float32的操作，先转换成int32再操作
func (this LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}

func (this LocalVars) GetFloat(index uint) float32 {
	bits := uint32(this[index].num)
	return math.Float32frombits(bits)
}

// SetLong 对long的操作,借助位运算
func (this LocalVars) SetLong(index uint, val int64) {
	this[index].num = int32(val)
	this[index+1].num = int32(val >> 32)
}

func (this LocalVars) GetLong(index uint) int64 {
	low := uint32(this[index].num)
	high := uint32(this[index+1].num)
	return int64(high)<<32 | int64(low)
}

// SetDouble 对double的处理,double->long ||然后按照long处理
func (this LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}

func (this LocalVars) GetDouble(index uint) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef 对引用值进行处理
func (this LocalVars) SetRef(index uint, ref *heap.Object) {
	this[index].ref = ref
}

func (this LocalVars) GetRef(index uint) *heap.Object {
	return this[index].ref
}

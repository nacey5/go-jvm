package runtime_data_area

import "math"

// OperandStack 操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

// 操作数栈的大小已经是编译确定了的
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		//初始化
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// 对int进行处理
func (this *OperandStack) PushInt(val int32) {
	this.slots[this.size].num = val
}
func (this *OperandStack) PopInt() int32 {
	this.size--
	return this.slots[this.size].num
}

// 对float进行处理
func (this *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	this.slots[this.size].num = int32(bits)
	this.size++
}
func (this *OperandStack) PopFloat() float32 {
	this.size--
	bits := uint32(this.slots[this.size].num)
	return math.Float32frombits(bits)
}

// 对long进行处理
func (this *OperandStack) PushLong(val int64) {
	this.slots[this.size].num = int32(val)
	this.slots[this.size+1].num = int32(val >> 32)
	this.size += 2
}
func (this *OperandStack) PopLong() int64 {
	this.size -= 2
	low := uint32(this.slots[this.size].num)
	high := uint32(this.slots[this.size+1].num)
	return int64(high)<<32 | int64(low)
}

// 对double进行处理
func (this *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	this.PushLong(int64(bits))
}
func (this *OperandStack) PopDouble() float64 {
	bits := uint64(this.PopLong())
	return math.Float64frombits(bits)
}

// 对引用类型进行处理
func (this *OperandStack) PushRef(ref *Object) {
	this.slots[this.size].ref = ref
	this.size++
}
func (this *OperandStack) PopRef() *Object {
	this.size--
	ref := this.slots[this.size].ref
	//将当前插槽引用的值置空,便于垃圾回收器进行回收
	this.slots[this.size].ref = nil
	return ref
}

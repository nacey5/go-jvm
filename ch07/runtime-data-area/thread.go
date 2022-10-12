package runtime_data_area

import "go-jvm/ch07/runtime-data-area/heap"

// Thread 线程结构体
type Thread struct {
	//pc机
	pc int
	//栈结构体
	stack *Stack
}

// NewThread 如果线程超过1024，stackoverflowException
// 如果可以动态扩展，内存不足，outOfMemoryException
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// Pc getter
func (this *Thread) Pc() int {
	return this.pc
}

// SetPc setter
func (this *Thread) SetPc(pc int) {
	this.pc = pc
}

func (this *Thread) PushFrame(frame *Frame) {
	this.stack.push(frame)
}

func (this *Thread) PopFrame() *Frame {
	return this.stack.pop()
}

func (this *Thread) CurrentFrame() *Frame {
	return this.stack.pop()
}

func (this *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(this, method)
}

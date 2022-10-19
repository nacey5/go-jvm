package runtime_data_area

import "go-jvm/ch10/runtime_data_area/heap"

/*
JVM

	Thread
	  pc
	  Stack
	    Frame
	      LocalVars
	      OperandStack
*/
type Thread struct {
	pc    int // the address of the instruction currently being executed
	stack *Stack
	// todo
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (this *Thread) PC() int {
	return this.pc
}
func (this *Thread) SetPC(pc int) {
	this.pc = pc
}

func (this *Thread) PushFrame(frame *Frame) {
	this.stack.push(frame)
}
func (this *Thread) PopFrame() *Frame {
	return this.stack.pop()
}

func (this *Thread) CurrentFrame() *Frame {
	return this.stack.top()
}
func (this *Thread) TopFrame() *Frame {
	return this.stack.top()
}

func (this *Thread) IsStackEmpty() bool {
	return this.stack.isEmpty()
}

func (this *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(this, method)
}

package runtime_data_area

import "go-jvm/ch08-2/runtime_data_area/heap"

// stack frame
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters & setters
func (this *Frame) LocalVars() LocalVars {
	return this.localVars
}
func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}
func (this *Frame) Thread() *Thread {
	return this.thread
}
func (this *Frame) Method() *heap.Method {
	return this.method
}
func (this *Frame) NextPC() int {
	return this.nextPC
}
func (this *Frame) SetNextPC(nextPC int) {
	this.nextPC = nextPC
}

func (this *Frame) RevertNextPC() {
	this.nextPC = this.thread.pc
}

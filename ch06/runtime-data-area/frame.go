package runtime_data_area

import "go-jvm/ch06/runtime-data-area/heap"

// lower用来实现链表结构
// localVars字段保存局部变量指针
// operandStack 保存操作数栈指针
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) SetThread(thread *Thread) {
	f.thread = thread
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f *Frame) Lower() *Frame {
	return f.lower
}

func (f *Frame) SetLower(lower *Frame) {
	f.lower = lower
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) SetLocalVars(localVars LocalVars) {
	f.localVars = localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) SetOperandStack(operandStack *OperandStack) {
	f.operandStack = operandStack
}

//func newFrame(maxLocals, maxStack uint) *Frame {
//
//}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (this *Frame) Method() *heap.Method {
	return this.method
}

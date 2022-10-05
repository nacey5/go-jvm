package runtime_data_area

// lower用来实现链表结构
// localVars字段保存局部变量指针
// operandStack 保存操作数栈指针
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
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

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

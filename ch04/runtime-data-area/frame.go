package runtime_data_area

// lower用来实现链表结构
// localVars字段保存局部变量指针
// operandStack 保存操作数栈指针
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
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

package runtime_data_area

type Stack struct {
	//保存栈的容量
	maxSize uint
	//保存栈当前的大小
	size uint
	//_top保存栈顶指针
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// push方法把帧推入栈顶
func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		panic("java.lang.StackOverFlowError")
	}
	if this._top != nil {
		frame.lower = this._top
	}
	this._top = frame
	this.size++
}

// 弹出栈帧
func (this *Stack) pop() *Frame {
	if this._top == nil {
		panic("jvm stack is empty")
	}
	top := this._top
	this._top = top.lower
	top.lower = nil
	this.size--
	return top
}

// top方法返回栈顶帧，但不弹出栈帧
func (this *Stack) top() *Frame {
	if this._top == nil {
		panic("jvm stack is empty")
	}
	return this._top
}

func (this *Stack) isEmpty() bool {
	return this._top == nil
}

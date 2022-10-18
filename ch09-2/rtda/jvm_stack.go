package rtda

// jvm stack
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // stack is implemented as linked list
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if this._top != nil {
		frame.lower = this._top
	}

	this._top = frame
	this.size++
}

func (this *Stack) pop() *Frame {
	if this._top == nil {
		panic("jvm stack is empty!")
	}

	top := this._top
	this._top = top.lower
	top.lower = nil
	this.size--

	return top
}

func (this *Stack) top() *Frame {
	if this._top == nil {
		panic("jvm stack is empty!")
	}

	return this._top
}

func (this *Stack) isEmpty() bool {
	return this._top == nil
}

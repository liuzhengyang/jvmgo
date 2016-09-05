package rtda

type Stack struct {
	maxSize uint
	size uint
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize:maxSize}
}

func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		panic("java.lang.StackOverFlowError")
	}

	if this._top != nil {
		frame.lower = this._top
	}
	this._top = frame
	this.size ++
}

func (this *Stack) pop() *Frame {
	if this._top == nil {
		panic("jvm stack is empty!")
	}

	top := this._top
	this._top = this._top.lower
	this.size --
	top.lower = nil
	return top
}

func (this *Stack) top() *Frame {
	if this._top == nil {
		panic("jvm stack is empty")
	}

	return this._top
}

